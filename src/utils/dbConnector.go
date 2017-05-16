package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db sql.DB

func InitDb() {
	Db, _ := sql.Open("mysql", "usimmonitor:uOslmBGUscQngIvu@tcp(10.9.0.3:3306)/usimmonitor")
	log.Println("Pinging Database")
	var err = Db.Ping()
	if err != nil {
		log.Fatal("Database connection failed")
	}
	log.Println("Database connection ok!")
}

func GetDb() *sql.DB {
	return &Db
}
