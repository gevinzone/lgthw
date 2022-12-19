package main

import (
	"github.com/gevinzone/lgthw/loop1/c07Micro/middleware"
	"log"
	"net/http"
	"os"
)

func main() {
	h := middleware.Handler
	h = middleware.ApplyMiddleware(h,
		middleware.Logger(log.New(os.Stdout, "", 0)),
		middleware.SetID(1000))
	http.HandleFunc("/", h)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
