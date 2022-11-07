package main

import "fmt"

func main() {

	chanOwner := func() <-chan int {
		// Channel is confined within this function. No other goroutine can
		// write to it.
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	// Consumer can only read.
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("received: %v\n", result)
		}
		fmt.Println("done")
	}

	results := chanOwner()
	consumer(results)
}
