package main

import "fmt"

func main() {
	vmap := map[int]foo{0: foo{value: 0}}
	fmt.Println("vmap:", vmap)
	//vmap[0].value = 3 // => ERROR

	pmap := map[int]*foo{0: &foo{value: 0}}
	fmt.Println("pmap[0]:", pmap[0])
	pmap[0].value = 3
	fmt.Println("pmap[0]:", pmap[0])
}

type foo struct {
	value int
}
