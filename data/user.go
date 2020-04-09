package data

type User struct {
	Acc string `json:"acc" gorm:"type:varchar(10);primary_key"`
	Pwd string `json:"pwd" gorm:"type:varchar(20)"`
}

type Tokens struct {
	Acc   string `json:"acc" gorm:"type:varchar(10)"`
	Token string `json:"token" gorm:"type:varchar(36);primary_key"`
}
