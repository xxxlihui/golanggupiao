package service

import (
	"math"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err.Error())
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

//分析当日的情况
func DayAnalyze(curDay *DayRecord) {
	if Float64Compare(curDay.PreClose, 0, 2) {
		//涨幅
		curDay.Zf = (curDay.Close - curDay.Open) / curDay.Open
		//涨停价
		curDay.Ztj = Float64Round(curDay.Open*1.1, 2)
		//跌停价
		curDay.Dtj = Float64Round(curDay.Open*0.9, 2)
		//涨停
		curDay.Zt = Float64Compare(curDay.Ztj, curDay.Close, 2)
		//跌停
		curDay.Dt = Float64Compare(curDay.Dtj, curDay.Close, 2)
		//大面
		curDay.Dm = (curDay.High-curDay.Close)/curDay.Close >= 0.1 || curDay.Dt
		//大肉
		curDay.Dr = (curDay.Close-curDay.Low)/curDay.Low >= 0.1 || curDay.Zt
		//破板
		curDay.Pb = Float64Compare(curDay.High, curDay.Ztj, 2) && !curDay.Zt
	} else {
		//涨幅
		curDay.Zf = (curDay.Close - curDay.PreClose) / curDay.Open
		//涨停价
		curDay.Ztj = Float64Round(curDay.PreClose*1.1, 2)
		//跌停价
		curDay.Dtj = Float64Round(curDay.PreClose*0.9, 2)
		//涨停
		curDay.Zt = Float64Compare(curDay.Ztj, curDay.Close, 2)
		//跌停
		curDay.Dt = Float64Compare(curDay.Dtj, curDay.Close, 2)
		//大面
		curDay.Dm = (curDay.High-curDay.Close)/curDay.Close >= 0.1 || curDay.Dt
		//大肉
		curDay.Dr = (curDay.Close-curDay.Low)/curDay.Low >= 0.1 || curDay.Zt
		//破板
		curDay.Pb = Float64Compare(curDay.High, curDay.Ztj, 2) && !curDay.Zt
	}
	//成交额大于20亿
	curDay.A20 = curDay.Amount >= 20*10000*10000
	//连板天数
	if curDay.Zt {
		curDay.Lb = 1
	}
	//是否是st
	curDay.St = strings.Index(curDay.Name, "st") > -1
	//是否是创业板
	curDay.Cy = curDay.Code >= 300000 && curDay.Code < 400000
}
