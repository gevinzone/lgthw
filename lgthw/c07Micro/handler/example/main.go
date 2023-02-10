package main

import (
	"github.com/gevinzone/lgthw/lgthw/c07Micro/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/name", handler.HelloHandler)
	http.HandleFunc("/greeting", handler.GreetingHandler)
	http.HandleFunc("/greeting/json", handler.GreetingJsonHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
