package main

import "fmt"

func main() {
	dataSource := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=true&loc=Local", "gevin", "gevin")
	db, err := Setup(dataSource)
	if err != nil {
		panic(err)
	}
	err = ExecWithTimeout(db)
	fmt.Println(err)
}
