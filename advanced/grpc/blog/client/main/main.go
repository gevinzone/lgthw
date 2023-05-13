package main

import (
	"github.com/gevinzone/lgthw/advanced/grpc/blog/client"
	"github.com/gevinzone/lgthw/advanced/grpc/blog/proto/gen"
	"google.golang.org/grpc"
)

func main() {
	target := ":8000"
	cc, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	c := gen.NewBlogAdminClient(cc)
	blogAdmin := client.NewBlogAdmin(c)
	err = blogAdmin.Start(":8080")
	if err != nil {
		panic(err)
	}
}
