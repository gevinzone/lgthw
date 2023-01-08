package math

import (
	"fmt"
	"math"
)

func Example() {
	i := 25
	fmt.Println(math.Sqrt(float64(i)))
	fmt.Println(math.Ceil(9.3))
	fmt.Println(math.Floor(9.9))
	fmt.Println("Pi: ", math.Pi, "E: ", math.E)
}
