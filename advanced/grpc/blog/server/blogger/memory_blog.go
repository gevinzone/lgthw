package blogger

import (
	"time"
)

type MemoryBlog struct {
}

func (m *MemoryBlog) CreateBlog(request *BlogRequest) (*Blog, error) {
	b := createDefaultBlog()
	return b, nil
}

func (m *MemoryBlog) GetBlog(id int64) *Blog {
	return createDefaultBlog()
}

func (m *MemoryBlog) UpdateBlog(request *BlogRequest) (*Blog, error) {
	return createDefaultBlog(), nil
}

func (m *MemoryBlog) DeleteBlog(id int64) error {
	return nil
}

func createDefaultBlog() *Blog {
	return &Blog{
		Id: 1,
		BlogRequest: BlogRequest{
			Title:   "default title",
			Content: "default content",
			Author:  "default author",
			Slug:    "default",
			Tags:    []string{"tag1", "tag2"},
		},
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
}
