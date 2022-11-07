package main

import (
	"fmt"
	"sync"
)

//
func main() {
	var c int
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			c = c + 1
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c)
}
