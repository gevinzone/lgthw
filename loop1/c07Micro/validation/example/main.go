package main

import (
	"github.com/gevinzone/lgthw/loop1/c07Micro/validation"
	"net/http"
)

func main() {
	h := validation.NewBaseHandler()
	http.HandleFunc("/", h.Process)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
