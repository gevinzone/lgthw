package ch8

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func SerialClock() error {
	listener, err := net.Listen("tcp", ":8001")
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		handleConn(conn)

	}
}

func ConcurrentClock() error {
	listener, err := net.Listen("tcp", ":8001")
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)

	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := conn.Write([]byte(time.Now().Format("15:04:05\n")))
		if err != nil {
			fmt.Println("error: ", err)
			return
		}
		time.Sleep(time.Second)
	}
}

func RunClient() error {
	conn, err := net.Dial("tcp", ":8001")
	if err != nil {
		return err
	}

	defer conn.Close()
	_, err = io.Copy(os.Stdout, conn)
	return err
}

func RunClient2(t time.Duration) error {
	conn, err := net.Dial("tcp", ":8001")
	if err != nil {
		return err
	}

	go func() {
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			if !errors.Is(err, net.ErrClosed) {
				fmt.Println("error: ", err)
			}
		}
	}()

	time.Sleep(t)
	return conn.Close()

}
