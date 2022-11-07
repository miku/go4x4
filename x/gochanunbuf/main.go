package main

import (
	"log"
	"time"
)

func main() {

	ch := make(chan int)
	go func() {
		log.Println("started consumer")
		for v := range ch {
			log.Println(v)
		}
	}()
	go func() {
		log.Println("started producer")
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	log.Println("main goroutine waits")
	time.Sleep(5 * time.Second)
}
