package service

import (
	"fmt"
	_ "github.com/lib/pq"
)
import "database/sql"

var db *sql.DB

func InitDb(host, user, password, dbname string, port int) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	_db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db = _db
}
func CloseDb() {
	db.Close()
}

func GetDB() *sql.DB {
	return db
}
