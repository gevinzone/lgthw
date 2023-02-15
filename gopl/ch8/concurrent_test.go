package ch8

import (
	"testing"
	"time"
)

func TestSerialClock(t *testing.T) {
	err := SerialClock()
	t.Log(err)
}

func TestConcurrentClock(t *testing.T) {
	err := ConcurrentClock()
	t.Log(err)
}

func TestRunClient(t *testing.T) {
	err := RunClient()
	t.Log(err)
}

func TestRunClient2(t *testing.T) {
	d := time.Second * 5
	err := RunClient2(d)
	t.Log(err)
}

func TestRunEchoServer(t *testing.T) {
	err := RunEchoServer()
	if err != nil {
		t.Error(err)
	}
}
