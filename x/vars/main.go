package main

import "log"

func main() {

	a := 3.14        // assign and infer type (only within functions)
	var b = 1.0      // assign and infer type (anywhere)
	var c int        // explicit type, without assignment
	var d int8 = 127 // explicit type, with assignment

	log.Println(a, b, c, d)
}
