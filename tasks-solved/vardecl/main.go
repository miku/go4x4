package main

import "fmt"

// Go has the "var" keyword, which can be used to declare variables, as well as a shorthand notation.
// Use 4 different styles of variable declaration and write a program that prints out:
//
// 0 b 123 d

func main() {

	var a int
	var b = "b"
	var c bool = true
	d := "d"

	fmt.Println(a, b, c, d) // 0 b true d
}
