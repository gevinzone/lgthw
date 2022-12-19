package handler

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGreetingHandler(t *testing.T) {
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
			GreetingHandler(tc.w, tc.r)
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
