package main

import (
	"database/sql"
	"fmt"
	//_ "github.com/go-sql-driver/mysql"
	"time"
)

type ExampleTable struct {
	Name    string
	Created time.Time
}

func Create(db *sql.DB) error {
	s := "CREATE TABLE example (name VARCHAR(20), created DATETIME)"
	if _, err := db.Exec(s); err != nil {
		return err
	}
	s = "INSERT INTO example (name, created) VALUES (?,?)"
	if _, err := db.Exec(s, "gevin", time.Now()); err != nil {
		return err
	}
	return nil
}

func Query(db *sql.DB) error {
	s := "SELECT name, created FROM example WHERE name=?"
	rows, err := db.Query(s, "gevin")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var e ExampleTable
		if err = rows.Scan(&e.Name, &e.Created); err != nil {
			return err
		}
		fmt.Printf("Results:\n\tName: %s\n\tCreated: %v\n", e.Name, e.Created)
	}
	return rows.Err()
}

func Exec(db *sql.DB) error {
	s := "DROP TABLE example"
	defer db.Exec(s)
	if err := Create(db); err != nil {
		return err
	}
	if err := Query(db); err != nil {
		return err
	}
	return nil
}
