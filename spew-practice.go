package main

import (
	"github.com/davecgh/go-spew/spew"
)

func main() {
	f := 1.23
	foo := Foo{bar: 9, baz: &f}
	spew.Dump(foo)
	spew.Dump(&foo)
}

type Foo struct {
	bar int
	baz *float64
}
