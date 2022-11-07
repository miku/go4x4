package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"
)

func main() {

	// By convention quit (or done) is passed first.
	newRandStream := func(quit <-chan bool) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited")
			defer close(randStream)
			for {
				select {
				case <-quit:
					return
				case randStream <- rand.Int():
				}
			}
		}()
		return randStream
	}

	quit := make(chan bool)
	randStream := newRandStream(quit)
	for i := 0; i < 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(quit)
	time.Sleep(100 * time.Microsecond)
	log.Println(runtime.NumGoroutine())

}
