package main

import (
	"github.com/gevinzone/lgthw/lgthw/c07Micro/handlerclass/closurehandler"
	"github.com/gevinzone/lgthw/lgthw/c07Micro/handlerclass/handler"
	"net/http"
)

func main() {
	h := handler.Handlers{}
	c := closurehandler.NewClosureHandlers(&closurehandler.MemStorage{})
	http.HandleFunc("/name", h.Get)
	http.HandleFunc("/greeting", h.Post)
	http.HandleFunc("/greeting/json", h.PostJson)
	http.HandleFunc("/get", c.Load(false))
	http.HandleFunc("/get/default", c.Load(true))
	http.HandleFunc("/set", c.Set())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
