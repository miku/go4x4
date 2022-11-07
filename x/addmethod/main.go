package main

import "fmt"

// We cannot define methods directly on int, but on a derived type.
// 	// https://stackoverflow.com/questions/28800672/how-to-add-new-methods-to-an-existing-type-in-go
type xint int

func (v xint) isEven() bool {
	return v%2 == 0
}

func main() {
	var x, y xint = 2, 3
	fmt.Printf("%v %v\n", x, x.isEven())
	fmt.Printf("%v %v\n", y, y.isEven())
}
