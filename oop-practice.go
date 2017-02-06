package main

import "fmt"

func main() {
	value := 3
	counter := Counter{Val: &value}
	fmt.Println(*counter.Val)
	counter.incr()
	fmt.Println(*counter.Val, value)
}

type Counter struct {
	Val *int
}

func NewCounter(container *int) *Counter {
	var c Counter
	c = Counter{Val: container}
	return &c
}

func (c *Counter) incr() {
	*(c.Val)++
}
