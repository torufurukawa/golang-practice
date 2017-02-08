package main

import "fmt"

// build(ch)().serve() で ch に送信

func main() {
	ch := make(chan bool)
	server := build(ch)()
	go server.Serve()
	_, ok := <-ch
	if !ok {
		fmt.Println("FAIL")
	}
	fmt.Println("OK")
}

func build(ch chan bool) func() TestServer {
	builder := func() TestServer {
		fmt.Println("building")
		return TestServer{ch: ch}
	}
	return builder
}

type TestServer struct {
	ch chan bool
}

func (s *TestServer) Serve() {
	fmt.Println("serving")
	s.ch <- true
}
