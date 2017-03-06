package main

import (
	"log"
)

func noopln(v ...interface{})               {}
func noopf(format string, v ...interface{}) {}

var Debugln func(v ...interface{}) = noopln
var Debugf func(format string, v ...interface{}) = noopf

func main() {
	Debugln("foo", "bar", 123)
	Debugf("%d %v", 123, map[int]string{1: "foo"})

	Debugln = log.Println
	Debugf = log.Printf

	Debugln("foo", "bar", 123)
	Debugf("%d %v", 123, map[int]string{1: "foo"})
}
