package main

import (
	"fmt"
	"log"
	"math"
)

const float64EqualityThreshold = 1e-11

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func Sqrt(x float64) float64 {
	var (
		z    = 1.0
		prev = 0.0
		i    = 0
	)
	for {
		i++
		z -= (z*z - x) / (2 * z)
		if almostEqual(prev, z) {
			break
		} else {
			prev = z
		}
	}
	log.Printf("num iterations: %d", i)
	return z
}

func main() {
	fmt.Println(Sqrt(2.0))
}
