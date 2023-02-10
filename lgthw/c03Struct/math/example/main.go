package main

import (
	"fmt"
	"github.com/gevinzone/lgthw/lgthw/c03Struct/math"
)

func main() {
	math.Example()

	for i := 0; i < 100; i++ {
		fmt.Println(math.Fib(i))
	}

}
