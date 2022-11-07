package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	N := 1000000
	var wg sync.WaitGroup
	wg.Add(N)
	started := time.Now()
	for i := 0; i < N; i++ {
		go func() {
			time.Sleep(150 * time.Millisecond)
			wg.Done()
		}()
	}
	wg.Wait()
	log.Printf("%d goroutines created in %v, %0.2f g/s", N, time.Since(started),
		float64(N)/float64(time.Since(started).Seconds()))
}
