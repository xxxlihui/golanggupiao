package data

import (
	"github.com/shopspring/decimal"
)

//日数据信息结构

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
	Open decimal.Decimal `json:"open" gorm:"type:numeric(20,2)"`
	//最高价
	High decimal.Decimal `json:"high"gorm:"type:numeric(20,2)"`
	//最低价
	Low decimal.Decimal `json:"low"gorm:"type:numeric(20,2)"`
	//收盘价
	Close decimal.Decimal `json:"close"gorm:"type:numeric(20,2)"`
	//成交量
	Volume decimal.Decimal `json:"volume"gorm:"type:numeric(20,2)"`
	//成交额
	Amount decimal.Decimal `json:"amount"gorm:"type:numeric(20,2)"`
	//上个交易日收盘价
	PreviousClose decimal.Decimal `json:"previousClose"gorm:"type:numeric(20,2)"`
	//涨停价
	LimitUp decimal.Decimal `json:"limitUp"gorm:"type:numeric(20,2)"`
	//跌停价
	LimitDown decimal.Decimal `json:"limitDown"gorm:"type:numeric(20,2)"`
}

//每一只票当日的一般的指标值
type PDaySample struct {
	//涨幅
	ChangePercent decimal.Decimal `json:"changePercent" gorm:"type:numeric(20,2)"`
	//涨幅
	Change decimal.Decimal `json:"change"gorm:"type:numeric(20,2)"`
	//振幅
	Amplitude decimal.Decimal `json:"amplitude"gorm:"type:numeric(20,2)"`
	//量比
	VolumeRate decimal.Decimal `json:"volumeRate"gorm:"type:numeric(20,2)"`
	//换手率
	TurnoverRate decimal.Decimal `json:"turnoverRate"gorm:"type:numeric(20,2)"`
	//市净率
	PB decimal.Decimal `json:"pb"gorm:"type:numeric(20,2)"`
	//市盈率（动态）
	PERation decimal.Decimal `json:"peRation"gorm:"type:numeric(20,2)"`
	//总市值
	TotalValue decimal.Decimal `json:"totalValue"gorm:"type:numeric(20,2)"`
	//流通市值
	CurrentValue decimal.Decimal `json:"currentValue"gorm:"type:numeric(20,2)"`
	//60涨跌幅
	Amplitude60 decimal.Decimal `json:"amplitude60"gorm:"type:numeric(20,2)"`
	//涨速
	Speed decimal.Decimal `json:"speed"gorm:"type:numeric(20,2)"`
	//5分钟涨跌
	Speed5 decimal.Decimal `json:"speed5"gorm:"type:numeric(20,2)"`
}

//每一次票当日的情绪
type PDayMood struct {
	//是否跌停
	IsLimitUp bool `json:"isLimitUp"`
	//是否涨停
	IsLimitDown bool `json:"isLimitDown"`
	//大肉
	Up10 bool `json:"up10"`
	//大面
	Down10 bool `json:"down10"`
	//破板
	FailLimitUp bool `json:"failLimitUp"`
	//连板天数
	LimitUpDays int `json:"limitUpDays"`
	//St股
	St bool `json:"st"`
	//创业板
	GEM bool `json:"gem"`
	//反脆弱
	ExceedExpect bool `json:"exceedExpect"`
	//反包 5日内重新涨停
	In5LimitUp bool `json:"in5LimitUp"`
	//一字板
	LimitUpOne bool `json:"limitUpOne"`
	//一字跌停
	LimitDownOne bool `json:"limitDownOne"`
}

//每日的复盘数据
type DayReplays struct {
	DayReplays []*DayReplay       `json:"dayReplays"`
	dayMapper  map[int]*DayReplay `json:"dayMapper"`
	//最早的数据日期
	MinDay int `json:"minDay"`
	//最晚的数据日期
	MaxDay int `json:"maxDay"`
	//加载的开始日期
	StartDay int `json:"startDay"`
	//加载的结束日期
	EndDay int `json:"endDay"`
}

//创建新的的默认的复盘数据信息表
func NewDefaultDayReplays() *DayReplays {
	return &DayReplays{dayMapper: map[int]*DayReplay{}}
}

func (receiver *DayReplays) Get(day int) *DayReplay {
	return receiver.dayMapper[day]
}

func (receiver *DayReplays) Init(replays []*DayReplay) {
	receiver.DayReplays = replays
	for _, replay := range replays {
		receiver.dayMapper[replay.Day] = replay
	}
}

//每日复盘
type DayReplay struct {
	PDay
	//大周期
	BigCycle int `json:"bigCycle"`
	//小周期
	SmallCycle int `json:"smallCycle"`
	//大肉个数
	Up10 int `json:"up10"`
	//大面个数
	Down10 int `json:"down10"`
	//回头波
	In1LimitUp int `json:"in1LimitUp"`
	//连板个数
	LimitUp int `json:"limitUp"`
	//炸板率
	FailLimitUp decimal.Decimal `json:"failLimitUp"`
	//一字板
	LimitUpOne int `json:"limitUpOne"`
	//昨日涨停溢价
	Premium decimal.Decimal `json:"premium"`
	//描述
	Description string `json:"description"`
	//标题
	Title string `json:"title"`
	//题材概念 排位
	Concepts []*Concept `json:"concepts"`
}

//题材概念
type Concept struct {
	//地位
	Order int `json:"order"`
	///涨停数量
	LimitUp int `json:"limitUp"`
	//拓展属性
	Expand string         `json:"expand"`
	Codes  []*ConceptCode `json:"codes"`
}
type ConceptCode struct {
	//是否是高度龙
	Top bool `json:"top"`
	//地位
	Order int `json:"order"`
	//代码
	Code string `json:"code"`
	//技术图形，说明
	Description []string `json:"description"`
}

//复盘数据存入数据库的格式
type DayReplayDb struct {
	Day  int `json:"day" gorm:"PRIMARY_KEY"`
	Data string
}
