package blogger

import (
	"github.com/gevinzone/lgthw/advanced/grpc/blog/proto/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type MemoryBlog struct {
}

func (m *MemoryBlog) CreateBlog(r *gen.BlogRequest) (*gen.Blog, error) {
	return createDefaultBlog(), nil
}

func (m *MemoryBlog) GetBlog(id int64) *gen.Blog {
	return createDefaultBlog()
}

func (m *MemoryBlog) UpdateBlog(r *gen.BlogRequest) (*gen.Blog, error) {
	return createDefaultBlog(), nil
}

func (m *MemoryBlog) DeleteBlog(id int64) error {
	return nil
}

func createDefaultBlog() *gen.Blog {
	slug := "default slug"
	return &gen.Blog{
		Id:         1,
		Title:      "default title",
		Content:    "default content",
		Author:     "default author",
		Slug:       &slug,
		Tags:       []string{"tag1", "tag2"},
		CreateTime: timestamppb.New(time.Now()),
		UpdateTime: timestamppb.New(time.Now()),
	}
}
