package main

import "fmt"
import "time"

func main() {
	done := false
	go do(&done)

	for {
		fmt.Println(done)
		if done == true {
			break
		}
		time.Sleep(time.Second * 1)
	}
}

func do(done *bool) {
	fmt.Println("will do")
	*done = true
	fmt.Println(done)
}

func oldmain() {
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
