package main

import "log" // why does it look like a string?

func main() {
	a := 1.0    // most used, only in funcs
	var b = 1.0 // sometimes
	var c int   // sometimes, "zero value"
	var d int8 = 127
	log.Println(a, b, c+1, d)

	s := "Hello 世界" // string
	log.Println(s)

	var e float32 = 1 // float32
	f := 1            // int
	var g byte = 97   // byte
	log.Println(e, f, g)
	log.Printf("%T %T %T", e, f, g)

	var arr [3]int
	log.Println(arr)

	var intSlice []int
	log.Println(intSlice, intSlice == nil) // []
	// data := []int{1, 2, 3, 4, 5}

	var s1 []string
	s2 := []string{}
	s3 := []string{"a", "b", "c"} // literal initialization
	log.Println(s1, s2, s3)

	var s4 []string
	s4 = append(s4, "x")
	s4 = append(s4, "y")
	s4 = append(s4, "z")

	s5 := []string{"X", "Y"}
	s4 = append(s4, s5...)
	log.Println(s4, s5)
	log.Println(len(s4), cap(s4))
	log.Println(len(s5), cap(s5))

	// log.Println(s5[2])
	// panic("stop it!")

	s6 := make([]int, 1000)
	// log.Println(s6, len(s6), cap(s6))
	for i := 0; i < len(s6); i++ {
		s6[i] = i
	}
	// log.Println(s6)
	s7 := s6[10:20]
	log.Println(s7, len(s7), cap(s7))

	s8 := []int{1, 2, 3}
	s9 := []int{0, 0}
	copy(s9, s8)
	log.Println(s8)
	log.Println(s9)
}
