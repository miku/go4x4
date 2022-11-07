package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

func longRunningProcess(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("cancelled")
			// teardown
			return
		default:
			fmt.Printf(".")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {

	ctx := context.Background()

	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	longRunningProcess(ctx)
}
