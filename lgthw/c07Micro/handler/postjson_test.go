package handler

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGreetingJsonHandler(t *testing.T) {
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
			GreetingJsonHandler(tc.w, tc.r)
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
