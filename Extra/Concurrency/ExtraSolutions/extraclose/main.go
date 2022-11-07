package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	data := make(chan int) // dummy channel

	var wg sync.WaitGroup
	var once sync.Once

	for i := 0; i < 24; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-data:
				log.Println("received data")
			case <-time.After(50 * time.Millisecond):
				log.Println("T")
				once.Do(func() {
					close(ch)
				})
			}
		}()
	}
	wg.Wait()
}
