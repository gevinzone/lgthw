package main

import (
	"context"
	"fmt"
	"github.com/gevinzone/lgthw/lgthw/c07Micro/grpc/proto/gen"
	"google.golang.org/grpc"
	"net"
)

type HelloServer struct {
	gen.UnimplementedGreeterServer
}

func (h *HelloServer) SayHello(ctx context.Context, request *gen.HelloRequest) (*gen.HelloReply, error) {
	reply := &gen.HelloReply{Message: "Hello, " + request.GetName()}
	return reply, nil
}

func main() {
	server := grpc.NewServer()
	gen.RegisterGreeterServer(server, &HelloServer{})
	address := ":8888"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Start serve on %s\n", address)
	server.Serve(listener)
}
