package db

import (
	"database/sql"
	"log"
)

type UserTable struct {
	dbHandle *sql.DB
}

func NewUserTable() *UserTable {
	return &UserTable{
		dbHandle: dbHandle,
	}
}

func (u *UserTable) IsExist(phone string) bool {
	rows, err := dbHandle.Query("SELECT * FROM users WHERE phone=?", phone)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return rows.Next()
}
