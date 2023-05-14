package grpc

import (
	"context"
	"github.com/gevinzone/lgthw/advanced/grpc/example/proto/gen"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	cc, err := grpc.Dial("register:///localhost:8001", grpc.WithInsecure(), grpc.WithResolvers(&ResolverBuilder{}))
	require.NoError(t, err)
	cli := gen.NewUserServiceClient(cc)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	resp, err := cli.GetById(ctx, &gen.GetByIdReq{Id: 1})
	require.NoError(t, err)
	t.Log(resp)
}
