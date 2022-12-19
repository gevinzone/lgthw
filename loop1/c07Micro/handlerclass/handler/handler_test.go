package handler

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var h Handlers

func TestHandlers_Get(t *testing.T) {
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
			h.Get(tc.w, tc.r)
			w := tc.w.(*httptest.ResponseRecorder)
			assert.Equal(t, tc.code, w.Code)
		})
	}
}
func TestHandlers_Post(t *testing.T) {
	testCases := []struct {
		name    string
		w       http.ResponseWriter
		r       *http.Request
		code    int
		wantRes GreetingResponse
	}{
		{
			name: "base case",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodPost, "/greeting?name=test&greeting=greeting", strings.NewReader(`{"name":"gevin","greeing":"hello"}`)),
			code: http.StatusOK,
			wantRes: GreetingResponse{
				Payload: struct {
					Greeting string `json:"greeting,omitempty"`
					Name     string `json:"name,omitempty"`
					Error    string `json:"error,omitempty"`
				}{
					Greeting: "greeting",
					Name:     "test",
				},
				Successful: true,
			},
		},
		{
			name: "bad method",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodGet, "/greeting?name=test&greeting=greeting", nil),
			code: http.StatusMethodNotAllowed,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			h.Post(tc.w, tc.r)
			w := tc.w.(*httptest.ResponseRecorder)
			assert.Equal(t, tc.code, w.Code)
			if w.Code != http.StatusOK {
				return
			}
			var res GreetingResponse
			err := json.Unmarshal(w.Body.Bytes(), &res)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
func TestHandlers_PostJson(t *testing.T) {
	testCases := []struct {
		name    string
		w       http.ResponseWriter
		r       *http.Request
		code    int
		wantRes GreetingResponse
	}{
		{
			name: "base case",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodPost, "/greeting", strings.NewReader(`{"name":"gevin","greeting":"hello"}`)),
			code: http.StatusOK,
			wantRes: GreetingResponse{
				Payload: struct {
					Greeting string `json:"greeting,omitempty"`
					Name     string `json:"name,omitempty"`
					Error    string `json:"error,omitempty"`
				}{
					Greeting: "hello",
					Name:     "gevin",
				},
				Successful: true,
			},
		},
		{
			name: "bad method",
			w:    httptest.NewRecorder(),
			r:    httptest.NewRequest(http.MethodGet, "/greeting", nil),
			code: http.StatusMethodNotAllowed,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			h.PostJson(tc.w, tc.r)
			w := tc.w.(*httptest.ResponseRecorder)
			assert.Equal(t, tc.code, w.Code)
			if w.Code != http.StatusOK {
				return
			}
			var res GreetingResponse
			err := json.Unmarshal(w.Body.Bytes(), &res)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
