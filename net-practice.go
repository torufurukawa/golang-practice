package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	go listen()

	// start echo on connection
	conn, err := net.Dial("tcp", "localhost:2000")
	if err != nil {
		log.Fatal(err)
	}
	// try echo several times
	fmt.Fprintf(conn, "hello\r\n")
	reader := bufio.NewReader(conn)
	status, err := reader.ReadString('\n')
	fmt.Println(status)
	fmt.Fprintf(conn, "hello\r\n")
	status, err = reader.ReadString('\n')
	fmt.Println(status)
	conn.Close()
	time.Sleep(time.Second * 5)
	log.Print("done")
}

func listen() {
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln("#", err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			defer func() { log.Println("####") }()
			reader := bufio.NewReader(c)
			for {
				status, err := reader.ReadString('\n')
				if err != nil {
					log.Println("$", err)
					return
				}
				fmt.Fprintf(c, status)
			}
		}(conn)
	}
}
