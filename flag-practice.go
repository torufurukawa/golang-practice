package main

import (
	"flag"
	"fmt"
)

func main() {
	foo := ""
	flag.StringVar(&foo, "foo", "DEFAULT", "foo parameter")
	flag.Parse()
	fmt.Println(foo)
}
