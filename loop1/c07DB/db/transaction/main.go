package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dataSource := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=true&loc=Local", "gevin", "gevin")
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// this won't do anything if commit is successful
	defer tx.Rollback()
	if err = Exec(db); err != nil {
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}
}
