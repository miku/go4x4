package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s1 []int
	fmt.Println(unsafe.Sizeof(s1)) // 24
}
