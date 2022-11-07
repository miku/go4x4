package main

import (
	"log"
	"time"
)

// Trying out the select statement.

// 1. Create a timer that fires after 1 second.

// 2. Create a goroutine that passes the values 1, 2, 3 on a int channel
// successively - with a 200ms sleep after each result.

// 3. Create a select statement (wrapped in a for loop) that has three cases:
// one for the timer, one for the channel sending int (1, 2, 3) and one for a
// `time.Tick` channel, that fires every e.g. 100ms.

func main() {

	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
			time.Sleep(200 * time.Millisecond)
		}
		// close(ch)
	}()

	timer := time.NewTimer(800 * time.Millisecond)

L:
	for {
		select {
		case <-time.Tick(100 * time.Millisecond):
			log.Println("tick")
		case <-timer.C:
			log.Println("timeout")
			break L
		case v, ok := <-ch:
			log.Println(v, ok)
		}
	}
}
