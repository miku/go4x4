package main

import "log"

// Integer is made up of all the int types
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Float is made up of all the float type
type Float interface {
	~float32 | ~float64
}

// Number is build from Integer and Float
type Number interface {
	Integer | Float
}

// Using Number
func Min[T Number](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func main() {
	// log.Println(Min(1, 2.0)) //  default type float64 of 2.0 does not match inferred type int for T
	log.Println(Min(1, 2))
	log.Println(Min(1.0, 2.0))
	log.Println(Min(int8(1), int8(2)))
}
