package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db sql.DB

func InitDb() {
	Db, err := sql.Open("mysql", "usimmonitor:uOslmBGUscQngIvu@tcp(10.9.0.3:3306)/usimmonitor")
	if err != nil {
		log.Fatal("Database connection failed")
	}
	log.Println("Database connection ok! %s", Db)
}

func GetDb() *sql.DB {
	return &Db
}
