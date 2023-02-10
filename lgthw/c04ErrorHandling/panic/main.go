package main

import (
	"fmt"
	"strconv"
)

func main() {
	PanicAndRecover()
}

func PanicAndRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic occurred: ", r)
			fmt.Println("recovered")
		}
	}()

	panicFunc := func() {
		zero, _ := strconv.ParseInt("0", 10, 64)
		_ = 1 / zero
		fmt.Println("we will never get here")
	}
	panicFunc()
}
