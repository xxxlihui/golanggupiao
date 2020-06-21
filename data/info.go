package data

type RecordInfo struct {
	Code         string `json:"code" gorm:"primary_key;varchar(8)"` //代码
	Name         string `json:"name" gorm:"varchar(20)"`            //名称
	LaunchDate   int    `json:"launchDate"`                         //上市日期
	DelistedDate int    `json:"delistedDate"`                       //退市日期

}
