package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan Event)
	go receive(ch)

	event := StartEvent{Val: 999}
	ch <- event
	time.Sleep(time.Second)
	close(ch)
}

type Event interface{}

type StartEvent struct {
	Val int
}

func receive(ch chan Event) {
	for {
		event, ok := <-ch
		if !ok {
			return
		}

		switch e := event.(type) {
		case StartEvent:
			fmt.Println(e.Val)
		}
	}
}
