package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("USER:", os.Getenv("USER"))

	wd, _ := os.Getwd()
	fmt.Println("wd:", wd)

	hostname, _ := os.Hostname()
	fmt.Println("hostname:", hostname)
}
