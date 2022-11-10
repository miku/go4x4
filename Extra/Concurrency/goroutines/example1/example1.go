// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to create goroutines and
// how the scheduler behaves.
package main

import (
	"fmt"
	"sync"
)

func main() {

	// wg is used to manage concurrency.
	var wg sync.WaitGroup // make the zero value useful
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// Create a goroutine from the lowercase function.
	go func() {
		defer wg.Done()
		lowercase()
	}()

	// Create a goroutine from the uppercase function.
	go func() {
		defer wg.Done()
		uppercase()
	}()

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

// lowercase displays the set of lowercase letters three times.
func lowercase() {

	// Display the alphabet three times
	for count := 0; count < 3; count++ {
		for r := 'a'; r <= 'z'; r++ {
			// time.Sleep(10 * time.Millisecond)
			fmt.Printf("%c ", r)
		}
	}
}

// uppercase displays the set of uppercase letters three times.
func uppercase() {

	// Display the alphabet three times
	for count := 0; count < 3; count++ {
		for r := 'A'; r <= 'Z'; r++ {
			// time.Sleep(10 * time.Millisecond)
			fmt.Printf("%c ", r)
		}
	}
}
