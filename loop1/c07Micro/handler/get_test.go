package handler

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	testCases := []struct {
		name string
		w    http.ResponseWriter
		r    *http.Request
		code int
	}{
		{
			name: "base-case",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodGet, "/name?name=test", nil),
			code: http.StatusOK,
		},
		{
			name: "bad method",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodPost, "/name?name=test", nil),
			code: http.StatusMethodNotAllowed,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			HelloHandler(tc.w, tc.r)
			w := tc.w.(*httptest.ResponseRecorder)
			assert.Equal(t, tc.code, w.Code)
		})
	}
}
