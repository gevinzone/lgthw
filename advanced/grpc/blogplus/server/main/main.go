package main

import (
	"github.com/gevinzone/lgthw/advanced/grpc/blogplus/server"
	"github.com/gevinzone/lgthw/advanced/grpc/blogplus/server/blogger"
)

func main() {
	admin := &blogger.MemoryBlog{}
	s := server.NewBlogServer(admin)
	err := s.Start(":8000")
	if err != nil {
		panic(err)
	}
}
