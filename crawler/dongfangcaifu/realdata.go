package dongfangcaifu

import (
	"encoding/json"
	"net/http"
	"nn/data"
	"nn/spider"
	"nn/util"
	"strings"
)

var rawUrl = "http://62.push2.eastmoney.com/api/qt/clist/get?cb=jQuery112403080608245235187_1588836679408&pn=1&pz=20000&po=1&np=1&ut=bd1d9ddb04089700cf9c27f6f7426281&fltt=2&invt=2&fid=f3&fs=m:0+t:6,m:0+t:13,m:0+t:80,m:1+t:2,m:1+t:23&fields=f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f13,f14,f15,f16,f17,f18,f20,f21,f23,f24,f25,f22,f11,f62,f128,f136,f115,f152&_=1588836679435"

func GetReal() ([]*data.RealData, error) {
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
		} `json:"_data"`
	}{}
	_json = strings.ReplaceAll(_json, `"-"`, "0")
	err = json.Unmarshal([]byte(_json), &_data)
	if err != nil {
		return nil, err
	}
	realDatas := make([]*data.RealData, len(_data.Data.Diff))
	for i := range _data.Data.Diff {
		realDatas[i] = toRealData(_data.Data.Diff[i])
	}
	return realDatas, nil
}

//实时数据
type Data struct {
	//代码
	Code string `json:"f12"`
	//名称
	Name string `json:"f14"`
	//最新价
	Close float64 `json:"f2"`
	//涨幅
	ChangePercent float64 `json:"f3"`
	//涨跌额
	Change float64 `json:"f4"`
	//成交量(手)
	Volume int64 `json:"f5"`
	//成交额
	Amount float64 `json:"f6"`
	//振幅
	Amplitude float64 `json:"f7"`
	//最高
	High float64 `json:"f15"`
	//最低
	Low float64 `json:"f16"`
	//今开
	Open float64 `json:"f17"`
	//昨收
	PreviousClose float64 `json:"f18"`
	//量比
	VolumeRate float64 `json:"f10"`
	//换手率
	TurnoverRate float64 `json:"f8"`
	//市净率
	PB float64 `json:"f23"`
	//市盈率(动态)
	PERation float64 `json:"f9"`
	//总市值
	TotalValue float64 `json:"f20"`
	//流通市值
	CurrentValue float64 `json:"f21"`
	//60日涨跌幅
	Amplitude60 float64 `json:"f24"`
	//年初至今涨跌幅
	Amplitude360 float64 `json:"f25"`
	//涨速
	Speed float64 `json:"f22"`
	//5分钟涨跌
	Speed5 float64 `json:"f11"`
}

func toRealData(_data *Data) *data.RealData {
	return &data.RealData{
		Code:          _data.Code,
		Name:          _data.Name,
		Close:         util.Float64TInt64(_data.Close, 2),
		ChangePercent: util.Float64TInt64(_data.ChangePercent, 2),
		Change:        util.Float64TInt64(_data.Change, 2),
		Volume:        _data.Volume,
		Amount:        util.Float64TInt64(_data.Amount, 2),
		Amplitude:     util.Float64TInt64(_data.Amplitude, 2),
		High:          util.Float64TInt64(_data.High, 2),
		Low:           util.Float64TInt64(_data.Low, 2),
		Open:          util.Float64TInt64(_data.Open, 2),
		PreviousClose: util.Float64TInt64(_data.PreviousClose, 2),
		VolumeRate:    util.Float64TInt64(_data.VolumeRate, 2),
		TurnoverRate:  util.Float64TInt64(_data.TurnoverRate, 2),
		PB:            util.Float64TInt64(_data.PB, 2),
		PERation:      util.Float64TInt64(_data.PERation, 2),
		TotalValue:    util.Float64TInt64(_data.TotalValue, 2),
		CurrentValue:  util.Float64TInt64(_data.CurrentValue, 2),
		Amplitude60:   util.Float64TInt64(_data.Amplitude60, 2),
		Amplitude360:  util.Float64TInt64(_data.Amplitude360, 2),
		Speed:         util.Float64TInt64(_data.Speed, 2),
		Speed5:        util.Float64TInt64(_data.Speed5, 2),
	}
}
