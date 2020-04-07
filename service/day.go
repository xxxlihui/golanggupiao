package service

//全部浮点数用int表示前三位表示小数位

type DayRecord struct {
	Day    int     `gorm:"PRIMARY_KEY";json:"day"`
	Open   int     `gorm:"PRIMARY_KEY";json:"open"`
	Code   int     `json:"code"`
	High   float32 `json:"high"`   //最高
	Low    float32 `json:"low"`    //最低
	Close  float32 `json:"close"`  //收盘
	Vol    uint64  `json:"vol"`    //成交量
	Amount float64 `json:"amount"` //成交额
	Zt     bool    `json:"zt"`     //是否涨停
	Dt     bool    `json:"dt"`     //是否跌停
	Zf     float32 `json:"zf"`     //涨幅
	Dm     bool    `json:"dm"`     //大面
	Dr     bool    `json:"dr"`     //大肉
	Pb     bool    `json:"pb"`     //破板
	A20    bool    `json:"a_20"`   //成交额大于20亿
	Lb     int     `json:"lb"`     //连板天数
	St     bool    `json:"st"`     //st股
	Cy     bool    `json:"cy"`     //创业板
}
type DayStat struct {
	Day    int     `gorm:"PRIMARY_KEY";json:"day"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Vol    uint64  `json:"vol"`
	Amount float64 `json:"amount"`
	Zt     int     `json:"zt"`   //涨停数量
	Dt     int     `json:"dt"`   //跌停数量
	Dm     int     `json:"dm"`   //大面数量
	Dr     int     `json:"dr"`   //大肉数量
	Pb     int     `json:"pb"`   //破板数量
	a20    int     `json:"a_20"` //成交额大于20亿数量
}
