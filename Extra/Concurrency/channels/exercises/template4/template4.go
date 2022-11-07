// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that creates a fixed set of workers to generate random
// numbers. Discard any number divisible by 2. Continue receiving until 100
// numbers are received. Tell the workers to shut down before terminating.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// Add imports.

func main() {

	rand.Seed(time.Now().UnixNano())

	// Create the channel for sharing results.
	values := make(chan int, 1)

	// Create a channel "shutdown" to tell goroutines when to terminate.
	shutdown := make(chan struct{})

	// Define the size of the worker pool. Use runtime.NumCPU to size the pool based on number of processors.
	poolSize := runtime.NumCPU()

	// Create a sync.WaitGroup to monitor the Goroutine pool. Add the count.
	var wg sync.WaitGroup
	wg.Add(poolSize)

	// Create a fixed size pool of goroutines to generate random numbers.
	for i := 0; i < poolSize; i++ {
		go func(id int) {
			defer wg.Done()
		OUT:
			// Start an infinite loop.
			for {

				// Generate a random number up to 1000.
				n := rand.Intn(1000)

				// Use a select to either send the number or receive the shutdown signal.
				select {

				// In one case send the random number.
				case values <- n:
					fmt.Printf("worker %d sent %d\n", id, n)

				// In another case receive from the shutdown channel.
				case <-shutdown:
					fmt.Printf("worker %d shutting down\n", id)
					break OUT
					// return
				}
			}
		}(i)
	}

	// Create a slice to hold the random numbers.
	var nums []int

	log.Printf("waiting ... (while %d goroutines are on)", runtime.NumGoroutine())
	time.Sleep(10 * time.Second)

	// Receive from the values channel with range.
	for i := range values {

		// continue the loop if the value was even.
		if i%2 == 0 {
			fmt.Printf("discarding %d\n", i)
			continue
		}

		// Store the odd number.
		nums = append(nums, i)

		// break the loop once we have 100 results.
		if len(nums) == 100 {
			break
		}
	}

	// Send the shutdown signal by closing the shutdown channel.
	close(shutdown)
	// Wait for the Goroutines to finish.
	wg.Wait()

	// Print the values in our slice.
	fmt.Println(nums)
}
