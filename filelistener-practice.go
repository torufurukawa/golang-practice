package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fin, err := os.Open("filelistener-dat.txt")
	fmt.Println("fin", fin)
	fmt.Println("error", err)

	listener, err := net.FileListener(fin)
	fmt.Println("listener", listener)
	fmt.Println("err", err)
}
