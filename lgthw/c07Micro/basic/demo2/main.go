package main

import (
	"net/http"
	"strings"
)

type handler struct {
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if strings.HasPrefix(r.URL.Path, "/other") {
		_, _ = w.Write([]byte("hello, others"))
		return
	}
	_, _ = w.Write([]byte("hello, world"))
}

func main() {
	h := &handler{}
	err := http.ListenAndServe(":8080", h)
	if err != nil {
		panic(err)
	}
}
