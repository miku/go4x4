package main

import (
	"fmt"
)

//
func main() {
	var c int
	for i := 0; i < 100; i++ {
		go func() {
			c = c + 1

		}()
	}
	fmt.Println(c)
}
