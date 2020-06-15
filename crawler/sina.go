package crawler

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"net/http"
	"nn/data"
	"nn/log"
	"nn/spider"
	"strconv"
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
				record := &data.DayRecord{}
				datas := strings.Split(dataStr, ",")
				record.Name = datas[0]
				v, _ := strconv.ParseFloat(datas[1], 64)
				record.Open = int(v * 100)
				v, _ = strconv.ParseFloat(datas[3], 64)
				record.Close = int(v * 100)
				v, _ = strconv.ParseFloat(datas[4], 64)
				record.High = int(v * 100)
				v, _ = strconv.ParseFloat(datas[5], 64)
				record.Low = int(v * 100)
				v, _ = strconv.ParseFloat(datas[9], 64)
				record.Amount = uint64(v * 100)
				u, _ := strconv.ParseUint(datas[8], 10, 64)
				record.Vol = u
				records = append(records, record)
			}
		} else {
			fmt.Printf("idx:%d\n", idx)
		}
	}
	return records, nil
}
