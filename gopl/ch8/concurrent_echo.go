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

type handleFunc func(net.Conn)

var msg = "Hello, World"
var msgSize = len([]rune(msg))

func RunEchoServer() error {
	return RunEchoServerBase(handleEchoConn)
}

func RunEchoServerE() error {
	return RunEchoServerBase(handleEchoConnE)
}

func RunEchoServerBase(handle handleFunc) error {
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
		go handle(conn)
	}
}

func handleEchoConn(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		echo(conn, scanner.Text(), time.Second)
	}
}

func handleEchoConnE(conn net.Conn) {
	defer conn.Close()
	// 这样没法在输入流中读数据
	//bs, err := io.ReadAll(conn)
	bs := make([]byte, msgSize)
	_, err := conn.Read(bs)
	if err != nil {
		fmt.Println("unexpected error: ", err)
		return
	}
	fmt.Println("received msg:", string(bs))
	echoE(conn, string(bs))

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

func echoE(conn net.Conn, text string) {
	if _, err := fmt.Fprintf(conn, "\t%s\n", strings.ToUpper(text)); err != nil {
		return
	}
	if _, err := fmt.Fprintf(conn, "\t%s\n", text); err != nil {
		return
	}
	if _, err := fmt.Fprintf(conn, "\t%s\n", strings.ToLower(text)); err != nil {
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

func RunEchoClientE() (err error) {
	var conn net.Conn
	if conn, err = net.Dial("tcp", ":8000"); err != nil {
		return err
	}
	_, err = conn.Write([]byte(msg))
	time.Sleep(time.Second * 2)
	bs := make([]byte, (msgSize+2)*3)
	_, err = conn.Read(bs)
	fmt.Println(string(bs))
	if err != nil {
		return err
	}
	return conn.Close()
}

func RunEchoClient2() (err error) {
	var conn net.Conn
	if conn, err = net.Dial("tcp", ":8000"); err != nil {
		return err
	}
	done := make(chan struct{})
	go func() {
		copyData(os.Stdout, conn)
		fmt.Println("done")
		done <- struct{}{}
	}()
	copyData(conn, os.Stdin)
	c := conn.(*net.TCPConn)
	//conn.Close()
	c.CloseWrite()
	<-done
	return nil
}

func copyData(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Println("unexpected error: ", err)
		return
	}
}
