package main

import "github.com/gevinzone/lgthw/gopl/ch8/chat"

func main() {
	addr := ":8000"
	err := chat.Connect(addr)
	if err != nil {
		panic(err)
	}
}
