package main

import "fmt"

func main() {

	var s = []string{"a", "b", "c", "d", "e"}

	s = append(s, "f")
	s = append(s[0:2], s[4:6]...)
	fmt.Println(s) // [a b e f]
}
