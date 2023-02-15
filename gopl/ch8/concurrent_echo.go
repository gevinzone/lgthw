package ch8

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

func RunEchoServer() error {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			if !errors.Is(err, net.ErrClosed) {
				fmt.Println("unexpected error: ", err)
			}
		}
		go handleEchoConn(conn)
	}
}

func handleEchoConn(conn net.Conn) {
	defer conn.Close()
	// 这样没法在输入流中读数据
	//bs, err := io.ReadAll(conn)
	//if err != nil {
	//	fmt.Println("unexpected error: ", err)
	//	return
	//}
	//if bs == nil {
	//	fmt.Println("no message found...")
	//	return
	//}
	//fmt.Println(string(bs))
	//echo(conn, string(bs), time.Second)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		echo(conn, scanner.Text(), time.Second)
	}
}

func echo(conn net.Conn, text string, t time.Duration) {
	// 这样输出数据可视化效果不好，格式化麻烦，改用fmt.Fprintln
	//if _, err := conn.Write([]byte(strings.ToUpper(text))); err != nil {
	//	return
	//}
	//time.Sleep(t)
	//if _, err := conn.Write([]byte(text)); err != nil {
	//	return
	//}
	//time.Sleep(t)
	//if _, err := conn.Write([]byte(strings.ToLower(text))); err != nil {
	//	return
	//}

	if _, err := fmt.Fprintln(conn, "\t", strings.ToUpper(text)); err != nil {
		return
	}
	time.Sleep(t)
	if _, err := fmt.Fprintln(conn, "\t", text); err != nil {
		return
	}
	time.Sleep(t)
	if _, err := fmt.Fprintln(conn, "\t", strings.ToLower(text)); err != nil {
		return
	}
}

func RunEchoClient() (err error) {
	var conn net.Conn
	if conn, err = net.Dial("tcp", ":8000"); err != nil {
		return err
	}
	go copyData(os.Stdout, conn)
	copyData(conn, os.Stdin)
	return conn.Close()
}

func copyData(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Println("unexpected error: ", err)
		return
	}
}
