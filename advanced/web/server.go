package web

import (
	"net"
	"net/http"
)

type HandleFunc func(ctx *Context)

type Server interface {
	http.Handler
	Start(addr string) error
	addRoute(method, path string, handler HandleFunc)
}

type HttpServer struct {
	router
}

var _ Server = &HttpServer{}

func (s *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &Context{
		Req:  r,
		Resp: w,
	}
	s.serve(ctx)
}

func (s *HttpServer) serve(ctx *Context) {
	n, ok := s.findRoute(ctx.Req.Method, ctx.Req.URL.Path)
	if !ok || n.handleFunc == nil {
		ctx.Resp.WriteHeader(http.StatusNotFound)
		_, _ = ctx.Resp.Write([]byte("Not Found"))
		return
	}
	n.handleFunc(ctx)
}

func (s *HttpServer) Get(path string, handler HandleFunc) {
	s.addRoute(http.MethodGet, path, handler)
}

func (s *HttpServer) Post(path string, handler HandleFunc) {
	s.addRoute(http.MethodPost, path, handler)
}

func (s *HttpServer) Put(path string, handler HandleFunc) {
	s.addRoute(http.MethodPut, path, handler)
}

func (s *HttpServer) Patch(path string, handler HandleFunc) {
	s.addRoute(http.MethodPatch, path, handler)
}

func (s *HttpServer) Delete(path string, handler HandleFunc) {
	s.addRoute(http.MethodDelete, path, handler)
}

func (s *HttpServer) Start(addr string) error {
	//err := http.ListenAndServe(addr, s)
	//return err
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	err = http.Serve(listener, s)
	return err
}
