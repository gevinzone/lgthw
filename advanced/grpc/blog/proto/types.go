package proto

import "github.com/gevinzone/lgthw/advanced/grpc/blog/proto/gen"

type Converter interface {
	ToBlog() *gen.Blog
}
