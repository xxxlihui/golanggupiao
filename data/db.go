package data

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB

func InitDb(host, user, password, dbname string, port int) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	fmt.Println(psqlInfo)
	_db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db = _db
	db.LogMode(true)
	db.AutoMigrate(&DayRecord{}, &DayRecord{},
		&User{}, &StockInfo{}, &Tokens{},
		&DayZt{}, //&data.Hudong{},
	)
}
func CloseDb() {
	db.Close()
}

func GetDB() *gorm.DB {
	return db
}
