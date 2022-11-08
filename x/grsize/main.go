package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}
	var c <-chan interface{}
	var wg sync.WaitGroup

	noop := func() {
		wg.Done()
		<-c // XXX
	}
	const numGoroutines = 1e5
	wg.Add(numGoroutines)

	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed() // XXX: Why the variety?

	fmt.Printf("%.3fkb\n", float64(after-before)/numGoroutines/1024)
	fmt.Printf("number of goroutines running: %v\n", runtime.NumGoroutine())
}
