package spider

import (
	"errors"
	"net/http/cookiejar"
	"nn/log"
	"sync"
	"time"
)

//基于cookie认证的处理函数，该函数在登陆成功后不再执行
type AuthenticationFunc interface {
	Auth(client Client) error
}

var (
	ErrLoginFail = errors.New("登陆失败")
	ErrNoStart   = errors.New("爬虫未开始")
)

const (
	//只需要登陆验证一次集可以，例如登陆后使用cookie
	AuthTypeOnce = iota + 1
	AuthTypeEvery
)
const LastTask = "xxxx-last-xxxx"

//要执行的队列
type Task struct {
	Name  string
	Fail  int16
	Try   int16
	Delay time.Duration
	Func  func(client *Client) error
}

//定义爬虫
type Spider struct {
	Client    *Client
	Auth      *Task
	AuthType  int
	AuthOk    bool
	Queue     chan *Task
	lock      sync.RWMutex
	start     bool
	threadNum int
	stop      chan struct{}
	loginLock sync.RWMutex
	async     bool
	pause     bool
	pauseChan chan struct{}
	loginChan chan struct{}
	DelayFunc func(*Task) time.Duration
}

func NewSpider(client *Client, authFunc *Task, threadNum int, delayFunc func(*Task) time.Duration) *Spider {
	spider := &Spider{Client: client, Auth: authFunc, threadNum: threadNum, DelayFunc: delayFunc}
	return spider
}

//只在一部执行的时候才能调用该函数
func (s *Spider) AddTask(task *Task) {
	if s.async && s.start {
		log.Debug("添加任务:%s", task.Name)
		s.Queue <- task
	} else {
		log.Debug("爬虫未开始不能添加任务")
	}

}
func (s *Spider) login() error {
	s.loginLock.Lock()
	defer s.loginLock.Unlock()

	log.Debug("执行登录")
	defer log.Debug("登录执行完成")
	if s.Auth == nil {
		return nil
	}
	if s.AuthOk {
		return nil
	}
	jar, _ := cookiejar.New(nil)
	s.Client.Client.Jar = jar
	s.Auth.Fail = 0
	if s.AuthType == AuthTypeOnce {
		for {
			if err := s.Auth.Func(s.Client); err != nil {
				if err == ErrLoginFail {
					return err
				}
				s.Auth.Fail++
				if s.Auth.Fail >= s.Auth.Try {
					return ErrLoginFail
				}
				time.Sleep(s.Auth.Delay * time.Second)
			}
			s.AuthOk = true
		}
	} else {
		s.AuthOk = true
		err := s.Auth.Func(s.Client)
		if err != nil {
			return err
		}
	}
	return nil
}
func (s *Spider) Pause() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.pause {
		return
	}
	s.pause = true
	s.pauseChan = make(chan struct{})
}
func (s *Spider) Resume() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !s.pause {
		return
	}
	s.pause = false
	close(s.pauseChan)
}
func (s *Spider) Stop() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.start == false {
		return
	}
	s.start = false
	s.async = false
	if s.async {
		close(s.stop)
		close(s.Queue)
	}
}
func (s *Spider) DoTask(task *Task) (_err error) {
	defer func() {
		if err := recover(); err != nil {
			_err = err.(error)
			log.Error("致命错误:%+s", _err.Error())
		}
	}()
	if !s.start {
		return ErrNoStart
	}
	if s.pause {
		<-s.pauseChan
	}
	if s.DelayFunc != nil {
		de := s.DelayFunc(task)
		log.Debug("延时执行[%s]%d秒", task.Name, de)
		time.Sleep(de * time.Second)
	}
	err := task.Func(s.Client)
	if err != nil {
		log.Error("任务[%s]执行失败:%s", task.Name, err.Error())
		if err == ErrLoginFail {
			s.AuthOk = false
			//记得defer是在函数完成的时候执行的不是在当前节点退回的时候执行的，所以这里会造成登录失败后程序一直卡住不执行
			return func() error {
				defer s.Resume()
				s.Pause()
				if err = s.login(); err != nil {
					log.Error("登陆失败，爬虫退出")
					s.Stop()
					return ErrLoginFail
				} else {
					s.Resume()
					return task.Func(s.Client)
				}
			}()
		} else {
			return err
		}
	} else {
		log.Info("任务[%s]执行成功", task.Name)
	}
	return nil
}
func (s *Spider) StartSync() error {
	return s._start(true)
}
func (s *Spider) StartAsync() error {
	return s._start(false)
}

//开始的并发的线程数
func (s *Spider) _start(issync bool) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.start {
		return nil
	}
	s.start = true
	if err := s.login(); err != nil {
		s.start = false
		return err
	}
	if issync {
		return nil
	}
	s.stop = make(chan struct{})
	s.Queue = make(chan *Task, 20)
	s.async = true
	for k := 0; k < s.threadNum; k++ {
		go func(_k int) {
			for {
				select {
				case <-s.stop:
					log.Debug("线程%d退出", _k)
					return
				case t := <-s.Queue:
					if s.pause {
						//暂停状态，先暂停for循环，暂停整个链条
						<-s.pauseChan
					}
					if t.Name == LastTask {
						return
					} else {

						s.DoTask(t)
					}
				}

			}

		}(k)
	}
	return nil
}
