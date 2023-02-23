package chat

import "testing"

func TestStartServer(t *testing.T) {
	addr := ":8000"
	err := StartServer(addr)
	if err != nil {
		t.Fatal(err)
	}
}

func TestConnect(t *testing.T) {
	addr := ":8000"
	err := Connect(addr)
	if err != nil {
		t.Fatal(err)
	}
}
