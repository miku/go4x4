package main

import "log"

// Write a function Rev for reversing a slice of any type.

func main() {
	vs := []int{3, 4, 5, 6, 7}
	log.Println(Rev(vs))
}

// $ go run rev.go
// 2022/11/07 22:04:28 [7 6 5 4 3]
