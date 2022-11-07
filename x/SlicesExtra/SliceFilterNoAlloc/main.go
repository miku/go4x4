package main

import "fmt"

func main() {
	var a = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceInfo("a", a)
	f := func(x int) bool { return x%2 == 0 }
	// alternative way to filter, reusing the underlying array; append w/o allocations
	b := a[:0]
	for _, x := range a {
		if f(x) {
			b = append(b, x)
			sliceInfo("b", b)
		}
	}
}

// [a] 0xc0000220f0 len=10 cap=10 [0 1 2 3 4 5 6 7 8 9]
// [b] 0xc0000220f0 len=1 cap=10 [0]
// [b] 0xc0000220f0 len=2 cap=10 [0 2]
// [b] 0xc0000220f0 len=3 cap=10 [0 2 4]
// [b] 0xc0000220f0 len=4 cap=10 [0 2 4 6]
// [b] 0xc0000220f0 len=5 cap=10 [0 2 4 6 8]

func sliceInfo(name string, s []int) {
	fmt.Printf("[%s] %p len=%d cap=%d %v\n", name, s, len(s), cap(s), s)
}
