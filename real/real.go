package real

import (
	"nn/chans"
	"nn/crawler/dongfangcaifu"
	"nn/data"
	"time"
)

var curCrawler = "东方财富"
var datas []*data.RealData
var lastMessage = ""

func Start() {
	go func() {
		_datas, err := dongfangcaifu.GetReal()
		if err != nil {
			lastMessage = err.Error()
		} else {
			lastMessage = ""
		}
		datas = _datas
		//实时消息通知
		chans.RealChan <- &struct {
			Message string
			Data    []*data.RealData
		}{Message: lastMessage, Data: _datas}
		if time.Now().Hour() > 15 {
			//收盘消息通知
			chans.EndChan<-_datas
			for {
				time.Sleep(1 * time.Second)
				if time.Now().Hour() > 9 && time.Now().Minute() > 15 {
					break
				}
			}
		}
		time.Sleep(3 * time.Second)
	}()
}
func GetReals() []*data.RealData {
	return datas
}
