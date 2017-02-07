package main

import "fmt"
import "errors"

func main() {
	value := 3
	counter := Counter{Val: &value}
	fmt.Println(*counter.Val)
	counter.incr()
	fmt.Println(*counter.Val, value)

	negValue := -10
	_, err := MakeCounter(negValue)
	fmt.Println(err)

	posValue := 99
	c, err := MakeCounter(posValue)
	fmt.Println(*c.Val)
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

func MakeCounter(val int) (*Counter, error) {
	if val < 0 {
		return nil, errors.New("negative")
	}

	counter := NewCounter(&val)
	return counter, nil
}
