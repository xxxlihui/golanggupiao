package service

import (
	"github.com/jinzhu/gorm"
	"math"
	"nn/data"
	"nn/log"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func checkDbError(err error) {
	if err != nil && gorm.ErrRecordNotFound != err {
		panic(err)
	}
}
func FloatCompare(f1, f2 float32, precision float64) bool {
	return Float64Compare(float64(f1), float64(f2), precision)
}
func Float64Round(f1, precision float64) float64 {
	return math.Round(f1*math.Pow(10, precision)) * math.Pow(0.1, precision)
}
func FloatRound(f1 float32, precision float64) float64 {
	return Float64Round(float64(f1), precision)
}
func Float64Compare(f1, f2 float64, precision float64) bool {
	n1 := math.Floor(math.Round(f1 * math.Pow(10, precision)))
	n2 := math.Floor(math.Round(f2 * math.Pow(10, precision)))
	return n1 == n2
}
func caseBool(b bool) int8 {
	if b {
		return 1
	}
	return 0
}

//分析当日的情况
func DayAnalyze(curDay *data.DayRecord) {
	if Float64Compare(curDay.PreClose, 0, 2) {
		//涨幅
		curDay.Zf = (curDay.Close - curDay.Open) / curDay.Open
		//涨停价
		curDay.Ztj = Float64Round(curDay.Open*1.1, 2)
		//跌停价
		curDay.Dtj = Float64Round(curDay.Open*0.9, 2)
		//涨停
		curDay.Zt = caseBool(Float64Compare(curDay.Ztj, curDay.Close, 2))
		//跌停
		curDay.Dt = caseBool(Float64Compare(curDay.Dtj, curDay.Close, 2))
		//大面
		curDay.Dm = caseBool((curDay.High-curDay.Close)/curDay.Close >= 0.1 || curDay.Dt == 1)
		//大肉
		curDay.Dr = caseBool((curDay.Close-curDay.Low)/curDay.Low >= 0.1 || curDay.Zt == 1)
		//破板
		curDay.Pb = caseBool(true == (Float64Compare(curDay.High, curDay.Ztj, 2) && curDay.Zt == 0))
	} else {
		//涨幅
		curDay.Zf = (curDay.Close - curDay.PreClose) / curDay.Open
		//涨停价
		curDay.Ztj = Float64Round(curDay.PreClose*1.1, 2)
		//跌停价
		curDay.Dtj = Float64Round(curDay.PreClose*0.9, 2)
		//涨停
		curDay.Zt = caseBool(Float64Compare(curDay.Ztj, curDay.Close, 2))
		//跌停
		curDay.Dt = caseBool(Float64Compare(curDay.Dtj, curDay.Close, 2))
		//大面
		curDay.Dm = caseBool(((curDay.High-curDay.Close)/curDay.Close >= 0.1 || curDay.Dt == 1))
		//大肉
		curDay.Dr = caseBool((curDay.Close-curDay.Low)/curDay.Low >= 0.1 || curDay.Zt == 1)
		//破板
		curDay.Pb = caseBool(Float64Compare(curDay.High, curDay.Ztj, 2) && curDay.Zt == 0)
	}
	//成交额大于20亿
	curDay.A20 = caseBool(curDay.Amount >= 20*10000*10000)
	//连板天数
	//这里无法分析
	if curDay.Zt == 1 {
		if curDay.Prelb > 0 {
			log.Debug("--------")
		}
		curDay.Lb = curDay.Prelb + 1
	} else {
		curDay.Lb = 0
	}
	if curDay.Zt == 1 && Float64Compare(curDay.High, curDay.Low, 2) {
		//一字涨停
		curDay.Ztyz = 1
	} else {
		curDay.Ztyz = 0
	}
	if curDay.Dt == 1 && Float64Compare(curDay.High, curDay.Low, 2) {
		//一字跌停
		curDay.Dtyz = 1
	} else {
		curDay.Dtyz = 0
	}
	//是否是st
	curDay.St = caseBool(strings.Index(curDay.Name, "st") > -1)
	//是否是创业板
	curDay.Cy = caseBool(strings.HasPrefix(curDay.Code, "300"))
}
