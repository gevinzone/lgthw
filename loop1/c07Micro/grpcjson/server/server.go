package main

import (
	"fmt"
	"github.com/gevinzone/lgthw/loop1/c07Micro/grpcjson/internal"
	"github.com/gevinzone/lgthw/loop1/c07Micro/grpcjson/keyvalue/gen"
	"google.golang.org/grpc"
	"net"
)

func main() {
	server := grpc.NewServer()
	gen.RegisterKeyValueServer(server, internal.NewKeyValue())
	address := ":8888"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Listening on port: %s\n", address)
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
