package main

import (
	"fmt"
	"time"
)

func main() {
	var data int
	go func() {
		data++
	}()

	// Note: Extensive testing found 5µs to be the ideal time to have chance to
	// observe different results on subsequent runs. Works on my machine.
	time.Sleep(5 * time.Microsecond)

	if data == 0 {
		fmt.Printf("value = %d\n", data)
	}
}
