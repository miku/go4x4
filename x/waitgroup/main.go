package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				fmt.Println(j)
				time.Sleep(10 * time.Millisecond)
			}
		}()
	}
	wg.Wait()
	fmt.Println("all done")
}
