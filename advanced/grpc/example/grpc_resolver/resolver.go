package grpc

import (
	"google.golang.org/grpc/resolver"
	"log"
)

type ResolverBuilder struct {
}

func (b *ResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &Resolver{cc: cc, target: target}
	r.ResolveNow(resolver.ResolveNowOptions{})
	return r, nil
}

func (b *ResolverBuilder) Scheme() string {
	return "register"
}

type Resolver struct {
	cc     resolver.ClientConn
	target resolver.Target
}

func (r *Resolver) ResolveNow(options resolver.ResolveNowOptions) {
	// 固定写死 IP 和端口
	// "localhost:8081"
	log.Print("url: ", r.target.URL.String())
	log.Print("endpoint: ", r.target.Endpoint())
	err := r.cc.UpdateState(resolver.State{
		Addresses: []resolver.Address{
			{
				Addr: "localhost:8000",
			},
		},
	})
	if err != nil {
		r.cc.ReportError(err)
	}
}

func (r *Resolver) Close() {

}
