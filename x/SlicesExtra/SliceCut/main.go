package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a = append(a[:2], a[7:]...) // [a] 0xc0000be000 len=4 cap=9 [1 2 8 9]
	sliceInfo("a", a)

	// the underlying array can accomodate this slice (and "5" becomes "visible" again)
	a = a[:5]
	sliceInfo("a", a)

	// a = a[:50] // panic: runtime error: slice bounds out of range [:50] with capacity 9
	// sliceInfo("a", a)
	// fmt.Println(a)
}

func sliceInfo(name string, s []int) {
	fmt.Printf("[%s] %p len=%d cap=%d %v\n", name, s, len(s), cap(s), s)
}

// [a] 0xc0000220f0 len=4 cap=9 [1 2 8 9]
// [a] 0xc0000220f0 len=5 cap=9 [1 2 8 9 5]
