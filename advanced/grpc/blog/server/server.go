package server

import (
	"github.com/gevinzone/lgthw/advanced/grpc/blog/proto/gen"
	"google.golang.org/grpc"
	"net"
)

type BlogServer struct {
	gen.UnimplementedBlogAdminServer
}

func (b *BlogServer) Start(address string) error {
	server := grpc.NewServer()
	gen.RegisterBlogAdminServer(server, b)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	return server.Serve(listener)
}
