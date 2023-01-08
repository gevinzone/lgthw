package main

import (
	"fmt"
	"github.com/gevinzone/lgthw/loop1/c03Struct/math"
)

func main() {
	math.Example()

	for i := 0; i < 100; i++ {
		fmt.Println(math.Fib(i))
	}

}
