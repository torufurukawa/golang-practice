package main

import (
	"bytes"

	"bufio"
	"fmt"
)

func main() {
	data := []byte("hello")
	reader := bufio.NewReader(bytes.NewReader(data))
	b := make([]byte, 3)
	reader.Read(b)
	fmt.Println(b)
}
