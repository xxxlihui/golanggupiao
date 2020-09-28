package crawler

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"net/http"
	"nn/data"
	"nn/log"
	"nn/spider"
	"strings"
)

func sinaGetByCodes(codes []string) ([]*data.DayRecord, error) {
	client := spider.NewClient(spider.RandomUserAgent())
	rspStr, _, _, err := spider.GetResponseString(func(bys []byte, head http.Header) string {
		by, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(bys)
		return string(by)
	}, func() (response *http.Response, e error) {
		return client.Get("http://hq.sinajs.cn/list="+intsJoin(codes),
			"www.sina.com", nil, nil)
	})
	if err != nil {
		log.Error("新浪接口发生错误,err:%s,返回内容:%s", err.Error(), rspStr)
		return nil, err
	}
	strs := strings.Split(rspStr, "\";\n")
	records := make([]*data.DayRecord, 0, len(codes))
	for _, e := range strs {
		if e == "" {
			continue
		}
		fmt.Println(e)
		idx := strings.Index(e, "hq_str_")
		if idx > -1 {
			code := e[idx+7+2 : idx+7+2+6]
			dataStr := e[idx+7+2+6+2:]
			fmt.Printf("code:%s,data:%s\n", code, dataStr)
			if dataStr != "" {
				//record := data.DayRecord{}

			}
		} else {
			fmt.Printf("idx:%d\n", idx)
		}
	}
	return records, nil
}
