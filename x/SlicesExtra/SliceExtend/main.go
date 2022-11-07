package main

import "fmt"

func main() {

	var a = []int{0, 1, 2, 3}
	sliceInfo("a", a)

	a = append(a, make([]int, 10)...)

	sliceInfo("a", a)

	a[13] = 13
	fmt.Printf("%v\n", a)
	// [a] 0xc00001c1c0 len=4 cap=4
	// [a] 0xc000062070 len=14 cap=14
	// [0 1 2 3 0 0 0 0 0 0 0 0 0 13]

}

func sliceInfo(name string, s []int) {
	fmt.Printf("[%s] %p len=%d cap=%d\n", name, s, len(s), cap(s))
}
