package chat

import (
	"bufio"
	"fmt"
	"net"
)

type client chan<- string

var (
	enter   = make(chan client)
	leaving = make(chan client)
	message = make(chan string)
)

func broadcast() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-message:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-enter:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	hello := who + " has arrived"
	fmt.Println(hello)
	message <- hello
	enter <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		//ch <- input.Text()
		message <- who + ": " + input.Text()
	}
	leaving <- ch
	bye := who + " has left"
	message <- bye
	fmt.Println(bye)
	_ = conn.Close()
}

func clientWriter(conn net.Conn, ch chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func StartServer(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	go broadcast()
	for {
		conn, er := listener.Accept()
		if er != nil {
			continue
		}
		go handleConn(conn)
	}
}
