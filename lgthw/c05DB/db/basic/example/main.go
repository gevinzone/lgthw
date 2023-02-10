package main

import (
	"fmt"
	"github.com/gevinzone/lgthw/lgthw/c05DB/db/basic"
)

func main() {
	dataSource := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=true&loc=Local", "gevin", "gevin")
	db, err := basic.Open(dataSource)
	if err != nil {
		panic(err)
	}
	err = basic.Exec(db)
	if err != nil {
		panic(err)
	}
}
