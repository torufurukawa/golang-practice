package main

import "fmt"

func main() {
	handlers := map[string]Handler{"PING": handlePing, "HELLO": handleHello}
	requests := []Request{
		Request{Cmd: "PING", Arg: ""},
		Request{Cmd: "HELLO", Arg: "wozozo"}}
	for _, req := range requests {
		handler := handlers[req.Cmd]
		resp := handler(req.Arg)
		fmt.Println(resp)
	}
}

func handlePing(arg string) string {
	return "PONG"
}

func handleHello(arg string) string {
	return "Bye, " + arg
}

type Handler func(string) string

type Request struct {
	Cmd string
	Arg string
}
