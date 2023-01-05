package main

import (
	"fmt"
	"github.com/gevinzone/lgthw/loop1/c06WebClient/rest/basicauth"
)

func main() {
	username, password := "gevin", "gevin"
	cli := basicauth.NewApiClient(username, password)
	url := "https://httpbin.org/basic-auth/gevin/gevin"
	code, err := cli.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(code)
}
