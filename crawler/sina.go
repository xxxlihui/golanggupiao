package crawler

import (
	"net/http"
	"nn/log"
	"nn/service"
	"nn/spider"
	"strings"
)

func sinaGetByCodes(codes []int) ([]*service.DayRecord, error) {
	client := spider.NewClient(spider.RandomUserAgent())
	rspStr, _, _, err := spider.GetResponseString(func() (response *http.Response, e error) {
		return client.Get("http://hq.sinajs.cn/list="+intsJoin(codes),
			"www.sina.com", nil, nil)
	})
	if err != nil {
		log.Error("新浪接口发生错误,err:%s,返回内容:%s", err.Error(), rspStr)
		return nil, err
	}
	strs:=strings.Split(rspStr,"\";\n")
	for e := range strs {
		code:=
	}


}
