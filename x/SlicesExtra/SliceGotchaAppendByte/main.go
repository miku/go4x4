package main

import (
	"fmt"
)

func sliceInfo(name string, s []byte) string {
	return fmt.Sprintf("[%s] %p len=%d cap=%d", name, s, len(s), cap(s))
}

func main() {
	a := []byte("Help")
	b := append(a, []byte(" Me")...)
	c := append(a, []byte(" Him")...)

	// WHAT GETS PRINTED HERE?
	fmt.Println("a", string(a))
	fmt.Println("b", string(b))
	fmt.Println("c", string(c))

	// fmt.Println("a", string(a), sliceInfo("a", a))
	// fmt.Println("b", string(b), sliceInfo("b", b))
	// fmt.Println("c", string(c), sliceInfo("c", c))
}
