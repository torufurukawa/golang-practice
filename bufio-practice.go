package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	// read 3 bytes
	data := []byte("hello")
	reader := bufio.NewReader(bytes.NewReader(data))
	b := make([]byte, 3)
	reader.Read(b)
	fmt.Println(b)
}
