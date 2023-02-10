package main

import (
	"github.com/gevinzone/lgthw/lgthw/c07Micro/validation"
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
