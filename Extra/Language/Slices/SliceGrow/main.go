package main

import "fmt"

func main() {

	var s []int
	for i := 0; i < 1000; i++ {
		s = append(s, 0)
		sliceInfo("s", s)
	}
}

func sliceInfo(name string, s []int) {
	fmt.Printf("[%s] %p len=%d cap=%d\n", name, s, len(s), cap(s))
}
