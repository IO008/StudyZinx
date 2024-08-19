package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var dbHandle *sql.DB = nil

func init() {
	dsn := "root:Mysqltest@tcp(127.0.0.1:3306)/chat"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	dbHandle = db
	fmt.Println("connected")
}

func GetDbHandle() *sql.DB {
	return dbHandle
}

func Close() {
	dbHandle.Close()
	dbHandle = nil
}
