package main

import "github.com/gevinzone/lgthw/advanced/grpc/blog/server"

func main() {
	s := server.BlogServer{}
	err := s.Start(":8000")
	if err != nil {
		panic(err)
	}
}
