package main

/*
#include <stdio.h>
#include <string.h>
double foo(char *name) {
	size_t length;
	length = strlen(name);
	return (double)length * 1.1;
}
*/
import "C"
import "fmt"

func main() {
	x := C.foo(C.CString("hello"))
	fmt.Println("x:", x)
}
