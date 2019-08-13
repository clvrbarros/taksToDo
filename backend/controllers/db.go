package controllers

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB
var err error

func init() {
	Database, err = sql.Open("sqlite3","database.db")
	Database.SetMaxOpenConns(3)
	if err != nil {
		fmt.Println(err)
	}
}