package main

import (
	"github.com/gevinzone/lgthw/lgthw/c07Micro/proxy"
	"net/http"
)

func main() {
	p := proxy.NewDefaultProxy("https://go.dev")
	http.Handle("/", p)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
