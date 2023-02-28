package basic

import "testing"

func TestServer(t *testing.T) {
	addr := ":8080"
	err := ListenAndServe(addr)
	t.Log(err)
}
