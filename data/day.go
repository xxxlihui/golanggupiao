package data

//全部浮点数用int表示前三位表示小数位

type DayRecord struct {
	Day      int    `gorm:"PRIMARY_KEY" json:"day"`
	Open     int    `json:"open"`
	Code     string `json:"code" gorm:"type:varchar(8);PRIMARY_KEY"`
	Name     string `json:"name" gorm:"type:varchar(5)"`
	High     int    `json:"high" `                      //最高
	Low      int    `json:"low" `                       //最低
	Close    int    `json:"close" `                     //收盘
	PreClose int    `json:"preClose" `                  //上一个交易日收盘价收盘
	Vol      uint64 `json:"vol"`                        //成交量
	Amount   uint64 `json:"amount"`                     //成交额
	Ztj      int    `json:"ztj" `                       //涨停价
	Zt       int8   `json:"zt" gorm:"type:smallint"`    //是否涨停
	Dt       int8   `json:"dt" gorm:"type:smallint"`    //是否跌停
	Dtj      int    `json:"dtj" `                       //跌停价                              //是否跌停
	Zf       int    `json:"zf"`                         //涨幅
	Zenf     int    `json:"zenf"`                       //振幅
	Dm       int8   `json:"dm"  gorm:"type:smallint"`   //大面
	Dr       int8   `json:"dr"  gorm:"type:smallint"`   //大肉
	Pb       int8   `json:"pb"  gorm:"type:smallint"`   //破板
	A20      int8   `json:"a_20"  gorm:"type:smallint"` //成交额大于20亿
	Lb       int    `json:"lb"`                         //连板天数
	Prelb    int    `json:"prelb"`                      //连板天数
	St       int8   `json:"st"  gorm:"type:smallint"`   //st股
	Cy       int8   `json:"cy"  gorm:"type:smallint"`   //创业板
	Fcr      int8   `json:"fcr"  gorm:"type:smallint"`  //反脆弱
	Fb       int8   `json:"fb"  gorm:"type:smallint"`   //反包
	Ztyz     int8   `json:"ztyz"  gorm:"type:smallint"` //一字板
	Tp       int8   `json:"tp"  gorm:"type:smallint"`   //突破
	Dtyz     int8   `json:"dzyz"  gorm:"type:smallint"` //跌停一字
}

//当日的统计
type DayStat struct {
	Day  int `gorm:"PRIMARY_KEY" json:"day"`
	Zt   int `json:"zt"`   //涨停数量
	Z    int `json:"z"`    //上涨数量
	D    int `json:"d"`    //下跌数量
	Dt   int `json:"dt"`   //跌停数量
	Dm   int `json:"dm"`   //大面数量
	Dr   int `json:"dr"`   //大肉数量
	Pb   int `json:"pb"`   //破板数量
	A20  int `json:"a_20"` //成交额大于20亿数量
	Mx   int `json:"mx"`   //最高板 高度
	Mx2  int `json:"mx2"`  //是否是次级最高板 高度
	Mxn  int `json:"mxn"`  //最高板的数量
	Mx2n int `json:"mx2n"` //次级最高板的数量
	Fb   int `json:"fb"`   //反包的数量
	Fcr  int `json:"fcr"`  //反脆弱的数量
	Tp   int `json:"tp"`   //突破的数量
	Ztyz int `json:"ztyz"` //一字涨停的数量
	Dtyz int `json:"dtyz"` //一字跌停的数量
}

//当日的连板统计
type DayZt struct {
	Day   int    `json:"day" gorm:"primary_key"` //日
	Zt    int    `json:"zt" gorm:"primary_key"`  //连板数
	Ztn   int    `json:"ztn"`                    //连板数量
	Mx    bool   `json:"mx"`                     //是否是最高板
	Mx2   bool   `json:"mx2"`                    //是否是次级最高板
	Codes string `json:"codes" gorm:"type:text"` //关联的代码
}

//分组
type Follows struct {
	Name string `json:"name" gorm:"type:text"`
}
