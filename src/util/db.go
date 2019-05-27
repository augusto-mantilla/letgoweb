package util

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB(dataSource string) {
	var err error
	db, err = sql.Open("mysql", dataSource)
	if err != nil {
		log.Panic(err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Panic(err.Error())
	}
}
