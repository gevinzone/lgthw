package middleware

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func ApplyMiddleware(h http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}

func Logger(l *log.Logger) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			l.Printf("started request to %s with id %s", r.URL, GetID(r.Context()))
			next(w, r)
			l.Printf("completed request to %s with id %s in %s", r.URL, GetID(r.Context()), time.Since(start))
		}
	}
}

func SetID(start int64) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), ID, strconv.FormatInt(start, 10))
			start++
			r = r.WithContext(ctx)
			next(w, r)
		}
	}
}
