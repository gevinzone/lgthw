package main

import "fmt"

func main() {
	DemoBasicError()
	err := DemoCustomError()
	fmt.Println("Demo custom error: ", err)
}
