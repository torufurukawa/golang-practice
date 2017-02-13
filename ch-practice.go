package main

import "fmt"
import "time"

func main() {
	done := make(chan bool)
	go handle(done)
	time.Sleep(time.Second * 1)
	close(done)
	time.Sleep(time.Second * 1)
}

func handle(done chan bool) {
	select {
	case <-done:
		fmt.Println("done")
		return
	}
}
