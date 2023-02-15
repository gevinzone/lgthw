package main

import (
	"github.com/gevinzone/lgthw/gopl/ch8"
)

//!+
func main() {
	err := ch8.RunEchoClient()
	if err != nil {
		panic(err)
	}
}
