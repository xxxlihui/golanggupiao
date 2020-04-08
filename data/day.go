package data

//全部浮点数用int表示前三位表示小数位

type DayRecord struct {
	Day      int     `gorm:"PRIMARY_KEY" json:"day"`
	Open     float64 `gorm:"PRIMARY_KEY" json:"open"`
	Code     string  `json:"code" gorm:"type:varchar(8)"`
	Name     string  `json:"name" gorm:"type:varchar(5)"`
	High     float64 `json:"high" gorm:"type:numeric(11,2)"`     //最高
	Low      float64 `json:"low" gorm:"type:numeric(11,2)"`      //最低
	Close    float64 `json:"close" gorm:"type:numeric(11,2)"`    //收盘
	PreClose float64 `json:"preClose" gorm:"type:numeric(11,2)"` //上一个交易日收盘价收盘
	Vol      uint64  `json:"vol"`                                //成交量
	Amount   float64 `json:"amount" gorm:"type:numeric(38,2)"`   //成交额
	Ztj      float64 `json:"ztj" gorm:"type:numeric(11,2)"`      //涨停价
	Zt       bool    `json:"zt"`                                 //是否涨停
	Dt       bool    `json:"dt"`                                 //是否跌停
	Dtj      float64 `json:"dtj" gorm:"type:numeric(11,2)"`      //跌停价                              //是否跌停
	Zf       float64 `json:"zf"`                                 //涨幅
	Dm       bool    `json:"dm"`                                 //大面
	Dr       bool    `json:"dr"`                                 //大肉
	Pb       bool    `json:"pb"`                                 //破板
	A20      bool    `json:"a_20"`                               //成交额大于20亿
	Lb       int     `json:"lb"`                                 //连板天数
	St       bool    `json:"st"`                                 //st股
	Cy       bool    `json:"cy"`                                 //创业板
}
type DayStat struct {
	Day    int     `gorm:"PRIMARY_KEY" json:"day"`
	High   float64 `json:"high" gorm:"type:decimal(11,2)"`
	Low    float64 `json:"low" gorm:"type:decimal(11,2)"`
	Close  float64 `json:"close" gorm:"type:decimal(11,2)"`
	Vol    uint64  `json:"vol" gorm:"type:decimal(11,2)"`
	Amount float64 `json:"amount" gorm:"type:decimal(38,2)"`
	Zt     int     `json:"zt"`   //涨停数量
	Dt     int     `json:"dt"`   //跌停数量
	Dm     int     `json:"dm"`   //大面数量
	Dr     int     `json:"dr"`   //大肉数量
	Pb     int     `json:"pb"`   //破板数量
	a20    int     `json:"a_20"` //成交额大于20亿数量
}

//分组
type Follows struct {
	Name string `json:"name" gorm:"type:text"`
}
