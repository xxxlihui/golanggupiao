package data

import (
	"github.com/shopspring/decimal"
)

type PDay struct {
	//日期20200203格式
	Day int `json:"day" gorm:"PRIMARY_KEY"`
}
type PCode struct {
	//股票代码sh600021
	Code string `json:"code" gorm:"type:varchar(8);PRIMARY_KEY"`
}

//日线数据
type PDayData struct {
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

//每一只票当日的一般的指标值
type PDaySample struct {
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

//每一次票当日的情绪
type PDayMood struct {
	//是否跌停
	IsLimitUp bool
	//是否涨停
	IsLimitDown bool
	//大肉
	Up10 bool
	//大面
	Down10 bool
	//破板
	FailLimitUp bool
	//连板天数
	LimitUpDays int
	//St股
	St bool
	//创业板
	GEM bool
	//反脆弱
	ExceedExpect bool
	//反包 5日内重新涨停
	In5LimitUp bool
	//一字板
	LimitUpOne bool
	//一字跌停
	LimitDownOne bool
}

//每日复盘
type DayReplay struct {
	PDay
	//大周期
	BigCycle int
	//小周期
	SmallCycle int
	//大肉个数
	Up10 int
	//大面个数
	Down10 int
	//回头波
	In1LimitUp int
	//连板个数
	LimitUp int
	//炸板率
	FailLimitUp decimal.Decimal
	//一字板
	LimitUpOne int
	//昨日涨停溢价
	Premium decimal.Decimal
	//描述
	Description string
	//标题
	Title string
	//题材概念 排位
	Concepts []*Concept
}

//题材概念
type Concept struct {
	//地位
	Order int
	///涨停数量
	LimitUp int
	//拓展属性
	Expand string
	Codes  []*ConceptCode
}
type ConceptCode struct {
	//是否是高度龙
	Top bool
	//地位
	Order int
	//代码
	Code string
	//技术图形，说明
	Description string
}
