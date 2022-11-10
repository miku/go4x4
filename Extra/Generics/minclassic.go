package main

import "log"

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	log.Println(MinInt(2, 3))
}
