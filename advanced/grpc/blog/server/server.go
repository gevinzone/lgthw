package server

import (
	"context"
	"github.com/gevinzone/lgthw/advanced/grpc/blog/proto/gen"
	"github.com/gevinzone/lgthw/advanced/grpc/blog/server/blogger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type BlogServer struct {
	gen.UnimplementedBlogAdminServer
	admin blogger.BlogAdmin
}

func NewBlogServer(a blogger.BlogAdmin) *BlogServer {
	b := &BlogServer{
		admin: a,
	}
	return b
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

func (b *BlogServer) GetArticle(ctx context.Context, id *gen.Id) (*gen.Blog, error) {
	blog := b.admin.GetBlog(id.GetId())
	return blog.ToBlog(), nil
}
func (b *BlogServer) CreateArticle(ctx context.Context, r *gen.BlogRequest) (*gen.Blog, error) {
	p := blogger.NewBlogRequest(blogger.WithRpcBlogRequest(r))
	blog, err := b.admin.CreateBlog(p)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return blog.ToBlog(), nil
}
func (b *BlogServer) UpdateArticle(ctx context.Context, r *gen.BlogRequest) (*gen.Blog, error) {
	p := blogger.NewBlogRequest(blogger.WithRpcBlogRequest(r))
	blog, err := b.admin.UpdateBlog(p)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return blog.ToBlog(), nil
}
func (b *BlogServer) DeleteArticle(ctx context.Context, id *gen.Id) (*gen.Response, error) {
	err := b.admin.DeleteBlog(id.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gen.Response{
		Code: 200,
	}, nil
}
