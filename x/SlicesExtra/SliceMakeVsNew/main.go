package main

import (
	"fmt"
)

func main() {
	v := new([]byte)
	w := make([]byte, 0)
	fmt.Printf("%T %v len=%v P %v V %v\n", v, v, len(*v), v == nil, *v == nil)
	fmt.Printf("%T %v len=%v V %v\n", w, w, len(w), w == nil)
	*v = append(*v, 1)
	fmt.Printf("%v %v %p", *v, v, *v) // %p
}
