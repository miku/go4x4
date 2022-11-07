package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		cap = 5
		ch  = make(chan string, cap) // a buffered channel
		num = 10
	)

	go func() {
		for p := range ch {
			fmt.Println("received from channel:", p)
		}
	}()

	for i := 0; i < num; i++ {
		select {
		case ch <- "data":
			fmt.Println("sent the data to the channel")
		default:
			fmt.Println("drop the data")
		}
	}

	time.Sleep(1 * time.Second)
	close(ch)
}
