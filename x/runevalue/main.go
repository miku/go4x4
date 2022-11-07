package main

import "fmt"

func main() {
	a := 'a'
	fmt.Printf("%T %c %d %x %v\n", a, a, a, a, a == 97)
	fmt.Printf("%s %d\n", "a", len("a"))

	v := '⧉'
	fmt.Printf("%T %c %d %x %v\n", v, v, v, v, v == 10697)

	// int32 a 97 61 true
	// int32 ⧉ 10697 29c9 true
}
