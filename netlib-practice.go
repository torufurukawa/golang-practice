package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.InterfaceAddrs())
	fmt.Println(net.Interfaces())
}
