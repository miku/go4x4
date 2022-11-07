package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func longRunning(ctx context.Context, wg *sync.WaitGroup) {
	defer func() {
		log.Println("exiting long running process")
		wg.Done()
	}()
	for {
		select {
		case <-ctx.Done():
			log.Println("cancelled")
			return
		default:
			log.Println("computing")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	var (
		wg            sync.WaitGroup
		ctx, cancelFn = context.WithCancel(context.Background())
	)
	log.Println("starting long running process")
	wg.Add(1)
	go longRunning(ctx, &wg)
	time.Sleep(2 * time.Second)
	cancelFn()
	log.Println("done")
	wg.Wait()
}
