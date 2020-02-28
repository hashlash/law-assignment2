package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(dataSourceString string) {
	var err error
	db, err = sql.Open("postgres", dataSourceString)
	if err != nil {
		fmt.Println(err)
	}

	if err = db.Ping(); err != nil {
		fmt.Println(err)
	}
}
