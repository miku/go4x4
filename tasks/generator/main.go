package main

import "fmt"

// Using channels to generate values.

// 1. Write a function that takes an int value "n" and returns a channal that
// can be used to retrieve numbers from 0 to n.

func main() {
	// Example usage could look like this.
	for i := range Seq(10) {
		fmt.Println(i)
	}
}
