package main

import (
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"net/url"
	"nn/spider"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	var url, dir string
	var threadCount int
	flags := []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "url",
			Aliases:     []string{"u", "URL"},
			Usage:       "地址",
			EnvVars:     []string{"URL"},
			Destination: &url,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "dir",
			Aliases:     []string{"d"},
			Usage:       "目录",
			EnvVars:     []string{"DIR"},
			Destination: &dir,
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:        "thread",
			Aliases:     []string{"t"},
			Usage:       "并发下载数，默认10",
			EnvVars:     []string{"THREAD"},
			Destination: &threadCount,
			Value:       10,
		}),
	}
	app := cli.App{
		Name:    "m38u下载器",
		Version: "1.0",
		Usage:   "m38u下载器,并自动合并",
		Authors: []*cli.Author{{Name: "lhn", Email: "550124023@qq.com"}},
		Flags:   flags,
		Before: altsrc.InitInputSourceWithContext(flags, func(context *cli.Context) (context2 altsrc.InputSourceContext, e error) {
			return altsrc.NewYamlSourceFromFile("set.yaml")
		}),
		Action: func(context *cli.Context) error {
			down(url, dir, threadCount)
			return nil
		},
	}
	app.Run(os.Args)
}

func down(_url, dir string, thread int) {
	list, err := getList(_url)
	if err != nil {
		println("下载错误", err.Error())
		return
	}
	l := strings.LastIndex(_url, "/")
	name := _url[l+1 : len(_url)-5]
	//pfx := _url[:l+1]
	total := len(list)
	tr := thread
	chans := make(chan *item, tr)
	okc := make(chan *item)
	//enc := make(chan struct{})
	for i := 0; i < tr; i++ {
		go func() {
			for {
				item := <-chans
				for {
					println("下载", item.name)
					u, _ := url.Parse(_url)
					u2, _ := u.Parse(item.name)
					err := downLoadOne(u2.String(), dir, name, item)
					if err != nil {
						println(item.name, "下载失败")
					} else {
						item.end(1)
						okc <- item
						break
					}
				}
			}
		}()
	}
	go func() {
		ok := 0
		for {
			i := <-okc
			ok++
			println("完成", i.name, "已完成", ok, "一共", total)
			if ok >= total {
				break
			}
		}
	}()
	os.Mkdir(filepath.Join(dir, name), os.ModePerm)

	for {
		ok := 0
		for _, v := range list {
			if v.status == 0 {
				v.start()
				chans <- v
				continue
			} else {
				//println("还有",v.name)
			}
			if v.status == 1 {
				ok++
			} else if v.status == 2 {
				//println("下载中", v.name)
			}
			//println("还有", total-ok)
		}
		time.Sleep(2 * time.Second)
		if ok >= total {
			break
		}
	}
	//合并文件
	merge(list, dir, name)
	os.RemoveAll(dir)
	println("下载完成")
}

func merge(items []*item, dir, name string) {
	p := filepath.Join(dir, name+".ts")
	f, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	defer f.Close()
	if err != nil {
		println("创建合并文件失败", err.Error())
		return
	}
	for _, v := range items {
		pth := filepath.Join(dir, name, v.Name())
		bys, err := ioutil.ReadFile(pth)
		if err != nil {
			println("读取子文件失败", err.Error())
			return
		}
		f.Write(bys)
	}
}

func write(bys []byte, pth string) error {
	f, err := os.OpenFile(pth, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		println("创建文件失败", err.Error())
		return err
	}
	defer f.Close()
	_, err = f.Write(bys)
	return err
}

func downLoadOne(url, dir, name string, item *item) error {
	pth := filepath.Join(dir, name, item.Name())
	info, err := os.Stat(pth)
	if err == nil {
		if info.Size() > 0 {

			return nil
		}
	}

	client := spider.NewClient(spider.RandomUserAgent())
	println("下载", url)
	bys, _, _, err := spider.GetResponseBytes(func() (*http.Response, error) {
		return client.Get(url, "", nil, nil)
	})
	if err != nil {
		println("下载", item.name, "失败")
		//item.status = 0
		return err
	}
	err = write(bys, pth)

	return err
}
func getTimeoutContext(second int) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(second)*time.Second)
	return ctx
}
func getList(url string) ([]*item, error) {
	client := spider.NewClient(spider.RandomUserAgent())
	str, _, _, err := spider.GetResponseString(nil, func() (*http.Response, error) {
		return client.Get(url, "", nil, nil)
	})
	if err != nil {
		return nil, err
	}
	strs := strings.Split(str, "\n")
	rstrs := make([]*item, 0)
	for _, v := range strs {
		if v == "" {
			continue
		}
		if !strings.HasPrefix(v, "#") {
			rstrs = append(rstrs, &item{name: v, status: 0})
		}
	}
	return rstrs, err
}

type item struct {
	name   string
	status int
	//lock   sync.Mutex
}

func (this *item) Name() string {
	idx := strings.LastIndex(this.name, "/")
	if idx >= 0 {
		return this.name[idx:]
	}
	return this.name
}
func (this *item) start() {
	//this.lock.Lock()
	this.status = 2
}
func (this *item) end(ok int) {
	this.status = ok
	//this.lock.Unlock()
}
func (this *item) rest() {
	this.status = 0
}
