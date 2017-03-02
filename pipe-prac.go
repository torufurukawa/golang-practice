package main

import (
	"fmt"
	"io"
)

func main() {
	reader, writer := io.Pipe()

	go func() {
		n, err := writer.Write([]byte("hello"))
		fmt.Println("Write() => ", n, err)
	}()

	buf := make([]byte, 10)
	n, err := reader.Read(buf)
	fmt.Println("Read() => ", n, err)
	fmt.Println("buf: ", string(buf))

	fmt.Println("hello")
}
