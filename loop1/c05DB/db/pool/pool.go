package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func Setup(datasource string) (*sql.DB, error) {
	db, err := sql.Open("mysql", datasource)
	if err != nil {
		return nil, err
	}
	// there will only ever be 24 open connections
	db.SetMaxOpenConns(24)
	// MaxIdleConns can never be less than MaxOpenConns, otherwise it'll default to that value
	db.SetMaxIdleConns(24)
	return db, nil
}

func ExecWithTimeout(db *sql.DB) error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now())
	defer cancel()
	_, err := db.BeginTx(ctx, nil)
	return err
}
