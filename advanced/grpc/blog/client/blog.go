package client

import (
	"context"
	"github.com/gevinzone/lgthw/advanced/grpc/blog/proto/gen"
	"net/http"
	"strconv"
	"strings"
)

type BlogAdmin struct {
	client gen.BlogAdminClient
}

func NewBlogAdmin(c gen.BlogAdminClient) *BlogAdmin {
	return &BlogAdmin{client: c}
}

func (b *BlogAdmin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/blog/get") {
		if r.Method != http.MethodGet {
			b.methodNotAllowed(w)
			return
		}
		b.getArticle(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/blog/create") {
		if r.Method != http.MethodPost {
			b.methodNotAllowed(w)
			return
		}
		b.getArticle(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/blog/update") {
		if r.Method != http.MethodPut {
			b.methodNotAllowed(w)
			return
		}
		b.getArticle(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/blog/delete") {
		if r.Method != http.MethodDelete {
			b.methodNotAllowed(w)
			return
		}
		b.getArticle(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Welcome to Blog Admin"))
}

func (b *BlogAdmin) getArticle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		b.writeError(w, err)
		return
	}
	blog, err := b.client.GetArticle(context.Background(), &gen.Id{Id: id})
	if err != nil {
		b.writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(blog.String()))
}

func (b *BlogAdmin) createArticle(w http.ResponseWriter, r *http.Request) {
	//id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	//if err != nil {
	//	b.writeError(w, err)
	//	return
	//}
	//blog, err := b.client.GetArticle(context.Background(), &gen.Id{Id: id})
	//if err != nil {
	//	b.writeError(w, err)
	//	return
	//}
	//w.WriteHeader(http.StatusOK)
	//_, _ = w.Write([]byte(blog.String()))
	panic("implement me")
}

func (b *BlogAdmin) updateArticle(w http.ResponseWriter, r *http.Request) {
	//id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	//if err != nil {
	//	b.writeError(w, err)
	//	return
	//}
	//blog, err := b.client.GetArticle(context.Background(), &gen.Id{Id: id})
	//if err != nil {
	//	b.writeError(w, err)
	//	return
	//}
	//w.WriteHeader(http.StatusOK)
	//_, _ = w.Write([]byte(blog.String()))
	panic("implement me")
}

func (b *BlogAdmin) deleteArticle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		b.writeError(w, err)
		return
	}
	res, err := b.client.DeleteArticle(context.Background(), &gen.Id{Id: id})
	if err != nil {
		b.writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(res.String()))
}

func (b *BlogAdmin) writeError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(err.Error()))
}

func (b *BlogAdmin) methodNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (b *BlogAdmin) Start(address string) error {
	return http.ListenAndServe(address, b)
}
