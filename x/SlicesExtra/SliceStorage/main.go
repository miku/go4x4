package main

import "fmt"

func main() {
	a := [3]int{0, 1, 2}        // array, literal
	var b [3]int                // array
	s := a[:]                   // from array to slice
	fmt.Printf("%T %T\n", a, s) // [3]int []int

	copy(b[:], s)  // copying from slice to array
	fmt.Println(b) // [0 1 2]

}
