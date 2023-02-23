package chat

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func Connect(addr string) error {
	conn, err := net.DialTimeout("tcp", addr, time.Second*3)
	if err != nil {
		return err
	}
	go copyData(conn, os.Stdin)
	copyData(os.Stdout, conn)

	return conn.Close()
}

func copyData(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		if err == net.ErrClosed {
			fmt.Println(err)
		} else {
			fmt.Println("unexpected error: ", err)
		}
		return
	}
}
