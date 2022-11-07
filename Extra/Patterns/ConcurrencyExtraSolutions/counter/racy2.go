package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//
func main() {
	var c int64
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			// c = c + 1
			atomic.AddInt64(&c, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c)
}
