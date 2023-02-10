package main

import (
	"flag"
	"fmt"
)

func main() {
	var name, key string
	flag.StringVar(&name, "name", "name", "name")
	flag.StringVar(&key, "key", "value", "key")
	flag.Parse()
	fmt.Println(name)
	fmt.Println(key)
	fmt.Println("Completed.")
}
