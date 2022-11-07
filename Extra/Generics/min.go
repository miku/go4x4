package main

import "log"

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

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

// 2022/11/07 22:16:31 1
// 2022/11/07 22:16:31 1
// 2022/11/07 22:16:31 1
