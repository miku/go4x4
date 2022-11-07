package main

import "fmt"

func main() {
	var a, b int
	fmt.Println(a == b)

	var aa, ba [3]byte
	fmt.Println(aa == ba)

	// var c float64
	// fmt.Println(a == c) //  invalid operation: a == c (mismatched types int and float64)

	// var as, bs []byte
	// fmt.Println(as == bs) // invalid operation: as == bs (slice can only be compared to nil)
}
