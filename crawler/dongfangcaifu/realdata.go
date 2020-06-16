package dongfangcaifu

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"net/http"
	"nn/data"
	"nn/spider"
	"strings"
)

var rawUrl = "http://62.push2.eastmoney.com/api/qt/clist/get?cb=jQuery112403080608245235187_1588836679408&pn=1&pz=20000&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f3&fs=m:0+t:6,m:0+t:13,m:0+t:80,m:1+t:2,m:1+t:23&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f11,f62,f128,f136,f115,f152&_=1588836679435"

func GetReal() ([]*RealData, error) {
	client := spider.NewClient(spider.RandomUserAgent())
	jsonp, _, _, err := spider.GetResponseString(nil, func() (*http.Response, error) {
		return client.Get(rawUrl, "", nil, nil)
	})
	if err != nil {
		return nil, err
	}
	start := strings.Index(jsonp, "(")
	end := strings.LastIndex(jsonp, ")")
	_json := jsonp[start+1 : end]
	_data := &struct {
		Data *struct {
			Diff []*Data `json:"diff"`
		} `json:"data"`
	}{}
	_json = strings.ReplaceAll(_json, `"-"`, "0")
	err = json.Unmarshal([]byte(_json), &_data)
	if err != nil {
		return nil, err
	}
	realData := make([]*RealData, len(_data.Data.Diff))
	for i := range _data.Data.Diff {
		realData[i] = toRealData(_data.Data.Diff[i])
	}
	return realData, nil
}

//实时数据
type Data struct {
	//代码
	Code string `json:"f12"`
	//名称
	Name string `json:"f14"`
	//最新价
	Close decimal.Decimal `json:"f2"`
	//涨幅
	ChangePercent decimal.Decimal `json:"f3"`
	//涨跌额
	Change decimal.Decimal `json:"f4"`
	//成交量(手)
	Volume decimal.Decimal `json:"f5"`
	//成交额
	Amount decimal.Decimal `json:"f6"`
	//振幅
	Amplitude decimal.Decimal `json:"f7"`
	//最高
	High decimal.Decimal `json:"f15"`
	//最低
	Low decimal.Decimal `json:"f16"`
	//今开
	Open decimal.Decimal `json:"f17"`
	//昨收
	PreviousClose decimal.Decimal `json:"f18"`
	//量比
	VolumeRate decimal.Decimal `json:"f10"`
	//换手率
	TurnoverRate decimal.Decimal `json:"f8"`
	//市净率
	PB decimal.Decimal `json:"f23"`
	//市盈率(动态)
	PERation decimal.Decimal `json:"f9"`
	//总市值
	TotalValue decimal.Decimal `json:"f20"`
	//流通市值
	CurrentValue decimal.Decimal `json:"f21"`
	//60日涨跌幅
	Amplitude60 decimal.Decimal `json:"f24"`
	//年初至今涨跌幅
	Amplitude360 decimal.Decimal `json:"f25"`
	//涨速
	Speed decimal.Decimal `json:"f22"`
	//5分钟涨跌
	Speed5 decimal.Decimal `json:"f11"`
}
type RealData struct {
	data.PCode
	data.PDayData
	data.PDaySample
}

func toRealData(_data *Data) *RealData {
	u := &RealData{}
	u.Code = _data.Code
	u.Open = _data.Open
	u.High = _data.High
	u.Low = _data.Low
	u.Close = _data.Close
	u.Volume = _data.Volume
	u.Amount = _data.Amount
	u.PreviousClose = _data.PreviousClose
	u.ChangePercent = _data.ChangePercent
	u.Amplitude = _data.Amplitude
	u.VolumeRate = _data.VolumeRate
	u.TurnoverRate = _data.TurnoverRate
	u.PERation = _data.PERation
	u.TotalValue = _data.TotalValue
	u.CurrentValue = _data.CurrentValue
	u.Amplitude60 = _data.Amplitude60
	u.Speed = _data.Speed
	u.Speed5 = _data.Speed5
	return u
}
