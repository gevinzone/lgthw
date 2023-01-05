package main

import (
	"fmt"
	"github.com/gevinzone/lgthw/loop1/c06WebClient/rest/client"
	"strings"
)

func main() {
	c := client.NewApiClient()
	var (
		url  string
		code int
		err  error
	)

	data := `{"key":"value"}`
	contentType := "application/json"
	//fmt.Println(data, contentType)

	url = "https://httpbin.org/get"
	code, err = c.Get(url)
	handleResponse("get", url, code, err)
	//
	url = "https://httpbin.org/post"
	code, err = c.Post(url, contentType, strings.NewReader(data))
	handleResponse("post", url, code, err)

	url = "https://httpbin.org/put"
	code, err = c.Put(url, contentType, strings.NewReader(data))
	handleResponse("put", url, code, err)

	url = "https://httpbin.org/delete"
	code, err = c.Delete(url)
	handleResponse("delete", url, code, err)
}

func handleResponse(method, url string, code int, err error) {
	if err != nil {
		panic(err)
	}
	fmt.Println(method, url, code)
}
