package main

import "fmt"
import "context"
import "time"
import "strings"

func main() {
	fmt.Println("hello")
	reqCh := make(chan string)
	respCh := make(chan string)
	complete := false
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ctx = context.WithValue(ctx, "reqCh", reqCh)
	ctx = context.WithValue(ctx, "respCh", respCh)
	ctx = context.WithValue(ctx, "complete", &complete)
	go serve(ctx)

	reqCh <- "foo"
	fmt.Println("resp:", <-respCh)
	fmt.Println("complete:", complete)
	cancel()
	time.Sleep(time.Millisecond * 100)
}

func serve(ctx context.Context) {
	fmt.Println("started serving")
	for {
		select {
		case s := <-ctx.Value("reqCh").(chan string):
			fmt.Println("req:", s)
			complete := ctx.Value("complete").(*bool)
			*complete = true
			ctx.Value("respCh").(chan string) <- strings.ToUpper(s)
		case <-ctx.Done():
			fmt.Println("done")
			return
		}
	}
}
