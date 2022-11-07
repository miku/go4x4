package main

import (
	"context"
	"log"
	"time"
)

func doSomething(ctx context.Context) {
	ctx, cancelCtx := context.WithCancel(ctx)

	log.Println("passed value was:", ctx.Value("key"))

	printCh := make(chan int)
	go doAnother(ctx, printCh)

	for num := 1; num <= 3; num++ {
		printCh <- num
	}

	cancelCtx()

	time.Sleep(100 * time.Millisecond)

	log.Printf("doSomething: finished\n")
}

func doAnother(ctx context.Context, printCh <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				log.Printf("doAnother err: %s\n", err)
			}
			log.Printf("doAnother: finished\n")
			return
		case num := <-printCh:
			log.Printf("doAnother: %d\n", num)
		}
	}
}

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "key", "value")

	doSomething(ctx)
}
