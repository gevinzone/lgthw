package blogger

import (
	"github.com/gevinzone/lgthw/advanced/grpc/blog/proto/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type BlogAdmin interface {
	CreateBlog(request *BlogRequest) (*Blog, error)
	GetBlog(id int64) *Blog
	UpdateBlog(request *BlogRequest) (*Blog, error)
	DeleteBlog(id int64) error
}

type BlogRequest struct {
	Title   string
	Content string
	Author  string
	Slug    string
	Tags    []string
}

func (b *BlogRequest) LoadFromRpc(r *gen.BlogRequest) *BlogRequest {
	b.Title = r.Title
	b.Tags = r.Tags
	b.Slug = *r.Slug
	b.Author = r.Author
	b.Content = r.Content
	return b
}

type BlogRequestOption func(request *BlogRequest)

func NewBlogRequest(opts ...BlogRequestOption) *BlogRequest {
	b := &BlogRequest{}
	for _, opt := range opts {
		opt(b)
	}
	return b
}

func WithRpcBlogRequest(r *gen.BlogRequest) BlogRequestOption {
	return func(b *BlogRequest) {
		b.Title = r.Title
		b.Tags = r.Tags
		b.Slug = *r.Slug
		b.Author = r.Author
		b.Content = r.Content
	}
}

type Blog struct {
	Id int64
	BlogRequest
	CreateTime time.Time
	UpdateTime time.Time
}

func (b *Blog) ToBlog() *gen.Blog {
	blog := &gen.Blog{
		Id: b.Id,
		Blog: &gen.BlogRequest{
			Title:   b.Title,
			Content: b.Content,
			Author:  b.Author,
			Slug:    &b.Slug,
			Tags:    b.Tags,
		},
		CreateTime: timestamppb.New(b.CreateTime),
		UpdateTime: timestamppb.New(b.UpdateTime),
	}
	return blog
}
