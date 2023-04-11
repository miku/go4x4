package main

import "log"

var Greeting = "Hello"

func main() {
	a := 3.14
	var b = 1.0
	var c int
	var d int8 = 127
	s := "hello 世界"
	var e int64 = int64(c)
	log.Println(a, b, c, d, e, s)

	var avs [4]int
	var vs []int
	data := []int{1, 2, 3, 4, 5} // literal init.

	log.Println(avs, vs, data)

	data = append(data, 6) // append an element
	log.Println(data)

	var aa []string
	aa = append(aa, "x")
	aa = append(aa, "y")

	bb := []string{"X", "Y"}
	aa = append(aa, bb...)
	log.Println(aa)

	log.Println(aa[0])
	log.Println(aa[0:2])

	// make([]int, 100) // type, length
	pre := make([]int, 0, 100) // type, length, capacity
	log.Println(pre, len(pre), cap(pre))

	stats := map[string]int{
		"ok":     120,
		"failed": 2,
	}
	log.Println(stats)

	stats["ok"]++
	log.Println(stats)
	stats["x"] = 1
	delete(stats, "x")
	log.Println(stats)

	// var dm map[string]int
	// dm["ok"] = 1 // panic
	mmap := make(map[string]int)
	log.Println(mmap)

	v := stats["bla"] // zero value
	log.Println(v)

	w, ok := stats["bla"]
	log.Println(w, ok)

}
