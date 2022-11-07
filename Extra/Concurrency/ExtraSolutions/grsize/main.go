package main

import (
	"log"
	"runtime"
	"sync"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		// ReadMemStats populates m with memory allocator statistics.
		// The returned memory allocator statistics are up to date as of the call to ReadMemStats.
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup

	noop := func() {
		wg.Done()
		<-c // A nil channel read blocks forever.
		log.Println("unreachable code")
	}

	log.Printf("%d goroutines before", runtime.NumGoroutine())

	const numGoroutines = 1000000
	wg.Add(numGoroutines)

	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	// All goroutines running.
	log.Printf("number of goroutines running: %v\n", runtime.NumGoroutine())

	after := memConsumed()

	log.Printf("%.3fkb\n", float64(after-before)/numGoroutines/1000)
}
