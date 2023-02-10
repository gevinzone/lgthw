package main

import (
	"context"
	"fmt"
	"github.com/gevinzone/lgthw/lgthw/c07Micro/grpc/proto/gen"
	"google.golang.org/grpc"
)

func main() {
	target := ":8888"
	cc, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := gen.NewGreeterClient(cc)
	in := &gen.HelloRequest{Name: "Gevin"}
	reply, err := client.SayHello(context.Background(), in)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)
}
