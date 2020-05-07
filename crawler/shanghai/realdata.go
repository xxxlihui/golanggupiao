package shanghai

import (
	"encoding/json"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"net/http"
	"nn/spider"
	"strings"
	"time"
)

//setcode 0 深圳 sz 1 上海 sh
var rawUrl = "http://yunhq.sse.com.cn:32041//v1/sh1/list/exchange/ashare?" +
	"callback=c&select=code%2Cname%2Copen%2Chigh%2Clow%2Clast%2Cprev_close%2C" +
	"chg_rate%2Cvolume%2Camount%2Ctradephase%2Cchange%2Camp_rate%2Ccpxxsubtype%2C" +
	"cpxxprodusta&order=&begin=0&end=10000&_={t}"

//获取上海证券交易所的实时交易数据
func GetRealData() ([]*RealData, error) {
	t := fmt.Sprint(time.Now().Unix())
	targetUrl := strings.Replace(rawUrl, "{t}", t, 1)

	client := spider.NewClient(spider.RandomUserAgent())
	rspStr, _, _, err := spider.GetResponseString(func(bys []byte, head http.Header) string {
		by, _ := simplifiedchinese.GBK.NewDecoder().Bytes(bys)
		return string(by)
	}, func() (response *http.Response, e error) {
		return client.Get(targetUrl,
			"", nil, nil)
	})
	if err != nil {
		return nil, err
	}
	start := strings.Index(rspStr, "(") + 1
	last := strings.LastIndex(rspStr, ")")
	str := rspStr[start:last]
	rspData := &RspData{}
	err = json.Unmarshal([]byte(str), &rspData)
	if err != nil {
		return nil, err
	}
	rd := make([]*RealData, len(rspData.List))
	for i, v := range rspData.List {
		rd[i] = &RealData{
			Code:         (v[0]).(string),
			Name:         v[1].(string),
			High:         v[2].(float64),
			Open:         v[3].(float64),
			Low:          v[4].(float64),
			Last:         v[5].(float64),
			PrevClose:    v[6].(float64),
			ChgClose:     v[7].(float64),
			Volume:       v[8].(float64),
			Amount:       v[9].(float64),
			Tradepahase:  v[10].(string),
			Change:       v[11].(float64),
			AmpRate:      v[12].(float64),
			Cpxxsubtype:  v[13].(string),
			Cpxxprodusta: v[14].(string),
		}
	}
	return rd, nil
}

type RspData struct {
	Date  int             `json:"date"`
	Time  int             `json:"time"`
	Total int             `json:"total"`
	Begin int             `json:"begin"`
	End   int             `json:"end"`
	List  [][]interface{} `json:"list"`
}

type RealData struct {
	Code         string  `json:"code"`
	Name         string  `json:"name"`
	Open         float64  `json:"open"`
	High         float64 `json:"high"`
	Low          float64 `json:"low"`
	Last         float64 `json:"last"`
	PrevClose    float64 `json:"prev_close"`
	ChgClose     float64 `json:"chg_close"`
	Volume       float64 `json:"volume"`
	Amount       float64 `json:"amount"`
	Tradepahase  string  `json:"tradepahase"`
	Change       float64 `json:"change"`
	AmpRate      float64 `json:"amp_rate"`
	Cpxxsubtype  string  `json:"cpxxsubtype"`
	Cpxxprodusta string  `json:"cpxxprodusta"`
}
