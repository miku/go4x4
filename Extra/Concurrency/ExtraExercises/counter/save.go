package main

import (
	"fmt"
)

// Exercise: Write a save counter.
//
// (1) Create a type counter that wraps an int and uses a sync.Mutex to protect access to the value.
// (2) Create a method on the type named `Inc` that increment the counter by one.
// (3) In main, create a counter, and start 100 goroutines, each incrementing the counter.
// (4) Print out the final value of the counter (by just accessing the struct field).
//

type counter struct{}

func main() {
	c := &counter{}
	for i := 0; i < 100; i++ {
		go func() {
			c.Inc()
		}()
	}

	fmt.Println(c.value)
}
