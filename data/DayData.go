package data

import (
	"github.com/shopspring/decimal"
)

//日线数据
type DayData struct {
	//日期20200203格式
	Day int `json:"day" gorm:"PRIMARY_KEY"`
	//股票代码sh600021
	Code string `json:"code" gorm:"type:varchar(8);PRIMARY_KEY"`
	//开盘价
	Open decimal.Decimal `json:"open" gorm:"type:decimal(11,2)"`
	//最高价
	High decimal.Decimal `json:"high gorm:"type:decimal(11,2)"`
	//最低价
	Low decimal.Decimal `json:"low gorm:"type:decimal(11,2)"`
	//收盘价
	Close decimal.Decimal `json:"close gorm:"type:decimal(11,2)"`
	//成交量
	Volume decimal.Decimal `json:"volume gorm:"type:decimal(11,2)"`
	//成交额
	Amount decimal.Decimal `json:"amount gorm:"type:decimal(11,2)"`
	//上个交易日收盘价
	PreviousClose decimal.Decimal `json:"previousClose gorm:"type:decimal(11,2)"`
	//涨停价
	LimitUp decimal.Decimal `json:"limitUp gorm:"type:decimal(11,2)"`
	//跌停价
	LimitDown decimal.Decimal `json:"limitDown gorm:"type:decimal(11,2)"`
}

//每一只票当日的情况
type DayState struct {
	//涨幅
	ChangePercent decimal.Decimal `json:"changePercent gorm:"type:decimal(11,2)"`
	//涨幅
	Change decimal.Decimal `json:"change gorm:"type:decimal(11,2)"`
	//振幅
	Amplitude decimal.Decimal `json:"amplitude gorm:"type:decimal(11,2)"`
	//量比
	VolumeRate decimal.Decimal `json:"volumeRate gorm:"type:decimal(11,2)"`
	//换手率
	TurnoverRate decimal.Decimal `json:"turnoverRate gorm:"type:decimal(11,2)"`
	//市净率
	PB decimal.Decimal `json:"pb gorm:"type:decimal(11,2)"`
	//市盈率（动态）
	PERation decimal.Decimal `json:"peRation gorm:"type:decimal(11,2)"`
	//总市值
	TotalValue decimal.Decimal `json:"totalValue gorm:"type:decimal(11,2)"`
	//流通市值
	CurrentValue decimal.Decimal `json:"currentValue gorm:"type:decimal(11,2)"`
	//60涨跌幅
	Amplitude60 decimal.Decimal `json:"amplitude60 gorm:"type:decimal(11,2)"`
	//涨速
	Speed decimal.Decimal `json:"speed gorm:"type:decimal(11,2)"`
	//5分钟涨跌
	Speed5 decimal.Decimal `json:"speed5 gorm:"type:decimal(11,2)"`
}
