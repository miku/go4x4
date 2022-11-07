package main

import "log"

func Rev[T any](vs []T) []T {
	rev := make([]T, len(vs))
	for i := 0; i < len(vs); i++ {
		rev[i] = vs[len(vs)-i-1]
	}
	return rev
}

func main() {
	vs := []int{3, 4, 5, 6, 7}
	log.Println(Rev(vs))
}

// $ go run rev.go
// 2022/11/07 22:04:28 [7 6 5 4 3]
