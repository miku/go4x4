package main

import "fmt"

func main() {
	var a = []int{0, 1, 2, 3}
	sliceInfo("a", a)
	i := 3 // insert position
	j := 4 // length
	//  --------------------------------------------------
	//                --------------------------------
	a = append(a[:i], append(make([]int, j), a[i:]...)...)
	sliceInfo("a", a)
	fmt.Println(a)
	// [a] 0xc00001c1c0 len=4 cap=4
	// [a] 0xc000020240 len=8 cap=8
	// [0 1 2 0 0 0 0 3]

}

func sliceInfo(name string, s []int) {
	fmt.Printf("[%s] %p len=%d cap=%d\n", name, s, len(s), cap(s))
}
