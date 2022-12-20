package main

import (
	"net/http"
	"strings"
)

type handler struct {
}

type Context struct {
	w http.ResponseWriter
	r *http.Request
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := Context{
		w: w,
		r: r,
	}
	// Do something before serve

	h.serve(ctx)

	// Do something after serve
}

func (h *handler) serve(ctx Context) {
	ctx.w.WriteHeader(http.StatusOK)
	if strings.HasPrefix(ctx.r.URL.Path, "/other") {
		_, _ = ctx.w.Write([]byte("hello, others"))
		return
	}
	_, _ = ctx.w.Write([]byte("hello, world"))
}

func (h *handler) Start(addr string) {
	err := http.ListenAndServe(addr, h)
	if err != nil {
		panic(err)
	}
}

func main() {
	h := &handler{}
	h.Start(":8080")
}
