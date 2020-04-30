package data

type Hudong struct {
	code string `json:"code" gorm:"primary_key;varchar(8)"` //代码
	url  string `json:"url" gorm:"varchar(360)"`
}
