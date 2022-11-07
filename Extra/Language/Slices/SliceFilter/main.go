package main

import "fmt"

func main() {
	var a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sliceInfo("a", a)
	// The filter function.
	keep := func(x int) bool {
		return x%2 == 0
	}

	// indexing filtered elements
	i := 0
	for _, x := range a {
		if keep(x) {
			a[i] = x
			i++
		}
	}
	sliceInfo("a", a)
	a = a[:i]
	sliceInfo("a", a)
	// fmt.Println(a[6]) // panic: runtime error: index out of range [6] with length 5
}

func sliceInfo(name string, s []int) {
	fmt.Printf("[%s] %p len=%d cap=%d %v\n", name, s, len(s), cap(s), s)
}

// [a] 0xc0000220a0 len=10 cap=10 [0 1 2 3 4 5 6 7 8 9]
// [a] 0xc0000220a0 len=10 cap=10 [0 2 4 6 8 5 6 7 8 9]
// [a] 0xc0000220a0 len=5 cap=10 [0 2 4 6 8]
