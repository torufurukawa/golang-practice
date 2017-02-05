package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("hello")
	reader := bufio.NewReader(bytes.NewReader(data))
	b := make([]byte, 3)
	reader.Read(b)
	fmt.Println(string(b)) // "hel"

	reader.Discard(1)

	b = make([]byte, 3)
	reader.Read(b)
	fmt.Println(string(b)) // "o"
}
