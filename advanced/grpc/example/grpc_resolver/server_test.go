package grpc

import (
	"context"
	"github.com/gevinzone/lgthw/advanced/grpc/example/proto/gen"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	us := &UserServer{}
	server := grpc.NewServer()
	gen.RegisterUserServiceServer(server, us)
	listener, err := net.Listen("tcp", ":8000")
	require.NoError(t, err)
	err = server.Serve(listener)
	require.NoError(t, err)
}

type UserServer struct {
	gen.UnimplementedUserServiceServer
}

func (u *UserServer) GetById(ctx context.Context, r *gen.GetByIdReq) (*gen.GetByIdResp, error) {
	return &gen.GetByIdResp{User: &gen.User{
		Id:   r.GetId(),
		Name: "Gevin Yu",
	}}, nil
}
