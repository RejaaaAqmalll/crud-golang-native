package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func KoneksiDB() {
	db, err := sql.Open("mysql", "root:@(localhost:3306)/db_golang?parseTime=true")
	if err != nil {
		panic(err)
	}

	log.Println("Database terkoneksi")
	DB = db
}
