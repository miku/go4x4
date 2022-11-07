package main

import (
	"log"

	"golang.org/x/exp/constraints"
)

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {

	log.Println(min("a", "b"))
}
