package web

import (
	"net"
	"net/http"
)

type HandleFunc func(ctx *Context)

type Server interface {
	http.Handler
	Start(addr string) error
	AddRoute(method, path string, handler HandleFunc)
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
	panic("implement me")
}

func (s *HttpServer) AddRoute(method, path string, handler HandleFunc) {
	s.addRoute(method, path, handler)
}

func (s *HttpServer) Get(path string, handler HandleFunc) {
	s.AddRoute(http.MethodGet, path, handler)
}

func (s *HttpServer) Post(path string, handler HandleFunc) {
	s.AddRoute(http.MethodPost, path, handler)
}

func (s *HttpServer) Put(path string, handler HandleFunc) {
	s.AddRoute(http.MethodPut, path, handler)
}

func (s *HttpServer) Patch(path string, handler HandleFunc) {
	s.AddRoute(http.MethodPatch, path, handler)
}

func (s *HttpServer) Delete(path string, handler HandleFunc) {
	s.AddRoute(http.MethodDelete, path, handler)
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
