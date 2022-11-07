package main

import "fmt"

// Union operator and type approximation both use together without interface
func Min[T ~int | ~float32 | ~float64](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// Type with underlying int
type Point int

func main() {
	// creating Point type
	x, y := Point(5), Point(2)

	fmt.Println(Min(x, y))

}
