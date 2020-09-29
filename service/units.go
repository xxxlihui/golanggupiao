package service

import (
	"github.com/jinzhu/gorm"
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

func caseBool(b bool) int8 {
	if b {
		return 1
	}
	return 0
}

//分析当日的情况
/*func DayAnalyze(curDay *data.DayRecord) {
	if curDay.PreClose == 0 {
		//涨幅
		curDay.Zf = util.FloatDiv(curDay.Close-curDay.Open, curDay.Open, 4)
		//振幅
		curDay.Zenf = util.FloatDiv(curDay.High-curDay.Close, curDay.Open, 4)
		//涨停价
		curDay.Ztj = util.FloatMul(curDay.Open, 110, 2)
		//跌停价
		curDay.Dtj = util.FloatMul(curDay.Open, 90, 2)
		//涨停
		curDay.Zt = caseBool(curDay.Ztj == curDay.Close)
		//跌停
		curDay.Dt = caseBool(curDay.Dtj == curDay.Close)
		//大面
		curDay.Dm = caseBool(util.FloatDiv(curDay.High-curDay.Close, curDay.Close, 2) >= 10 || curDay.Dt == 1)
		//大肉
		curDay.Dr = caseBool(util.FloatDiv(curDay.Close-curDay.Low, curDay.Low, 2) >= 10 || curDay.Zt == 1)
		//破板
		curDay.Pb = caseBool(curDay.High == curDay.Ztj && curDay.Zt == 0)
	} else {
		//涨幅
		curDay.Zf = util.FloatDiv(curDay.Close-curDay.PreClose, curDay.PreClose, 4)
		//振幅
		curDay.Zenf = util.FloatDiv(curDay.High-curDay.Low, curDay.PreClose, 4)
		//涨停价
		curDay.Ztj = util.FloatMul(curDay.PreClose, 110, 2)
		//跌停价
		curDay.Dtj = util.FloatMul(curDay.PreClose, 90, 2)
		//涨停
		curDay.Zt = caseBool(curDay.Ztj == curDay.Close)
		//跌停
		curDay.Dt = caseBool(curDay.Dtj == curDay.Close)
		//大面
		curDay.Dm = caseBool(util.FloatDiv(curDay.High-curDay.Close, curDay.Close, 2) >= 10 || curDay.Dt == 1)
		//大肉
		curDay.Dr = caseBool(util.FloatDiv(curDay.Close-curDay.Low, curDay.Low, 2) >= 10 || curDay.Zt == 1)
		//破板
		curDay.Pb = caseBool(curDay.High == curDay.Ztj && curDay.Zt == 0)
	}
	//成交额大于20亿
	curDay.A20 = caseBool(curDay.Amount >= 20*10000*10000*100)
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
	if curDay.Zt == 1 && curDay.High == curDay.Low {
		//一字涨停
		curDay.Ztyz = 1
	} else {
		curDay.Ztyz = 0
	}
	if curDay.Dt == 1 && curDay.High == curDay.Low {
		//一字跌停
		curDay.Dtyz = 1
	} else {
		curDay.Dtyz = 0
	}
	//是否是st
	curDay.St = caseBool(strings.Index(curDay.Name, "st") > -1)
	//是否是创业板
	curDay.Cy = caseBool(strings.HasPrefix(curDay.Code, "300"))
}*/
