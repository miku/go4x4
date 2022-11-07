package main

import (
	"log"
	"time"
)

func main() {
	c := make(chan string, 1)
	go func() {
		log.Println("long running function")
		time.Sleep(200 * time.Millisecond)
		c <- "ok"
	}()
	select {
	case v := <-c:
		log.Printf("received: %s", v)
	case <-time.After(50 * time.Millisecond):
		log.Println("timed out")
	}
}
