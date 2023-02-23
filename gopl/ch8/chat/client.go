package chat

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func Connect(addr string) error {
	conn, err := net.DialTimeout("tcp", addr, time.Second*3)
	if err != nil {
		return err
	}
	defer conn.Close()
	done := make(chan struct{})
	go func() {
		_, _ = io.Copy(conn, os.Stdin)
		done <- struct{}{}
	}()
	mustCopyData(os.Stdout, conn)
	<-done

	return nil
}

func mustCopyData(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
