package blogger

import "github.com/gevinzone/lgthw/advanced/grpc/blogplus/proto/gen"

type BlogAdmin interface {
	CreateBlog(request *gen.BlogRequest) (*gen.Blog, error)
	GetBlog(id int64) *gen.Blog
	UpdateBlog(request *gen.BlogRequest) (*gen.Blog, error)
	DeleteBlog(id int64) error
}
