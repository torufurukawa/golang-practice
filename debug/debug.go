package debug

import (
	"log"
)

func noopln(v ...interface{})               {}
func noopf(format string, v ...interface{}) {}

var Println func(v ...interface{}) = noopln
var Printf func(format string, v ...interface{}) = noopf

func Enable() {
	Println = log.Println
	Printf = log.Printf
}

func Disable() {
	Println = noopln
	Printf = noopf
}
