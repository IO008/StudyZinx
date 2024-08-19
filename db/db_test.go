package db

import (
	"fmt"
	"log"
	"testing"
	"time"
	//_ "github.com/go-sql-driver/mysql"
)

func TestDB(t *testing.T) {

	db := GetDbHandle()
	defer Close()

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected")

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var phone string
		var first_name string
		var last_name string
		err := rows.Scan(&phone, &first_name, &last_name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("phone: %s, first_name: %s, last_name: %s\n", phone, first_name, last_name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)
}
