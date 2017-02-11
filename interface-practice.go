package main

import "fmt"

func main() {
	invoke(NewMyRunner())
}

func handle(conn bool) {
	fmt.Println("handle")
}

type Server interface {
	Serve()
}

type Runner interface {
	Run()
}

type MyServer struct{}

func NewMyServer() MyServer {
	return MyServer{}
}

func (s *MyServer) Serve() {
	fmt.Println("Hello")
}

type MyRunner struct {
	server MyServer
}

func NewMyRunner() MyRunner {
	return MyRunner{server: NewMyServer()}
}

func (r MyRunner) Run() {
	r.server.Serve()
}

func invoke(runner Runner) {
	runner.Run()
}
