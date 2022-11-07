package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
)

func main() {

	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited")
			defer close(randStream)
			for {
				randStream <- rand.Int()
			}
		}()
		return randStream
	}

	randStream := newRandStream()
	for i := 0; i < 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	log.Println(runtime.NumGoroutine())
}
