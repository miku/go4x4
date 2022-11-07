package main

import "log"

// Predefined types do not implement any method, so the following won't work.

type Sixteen interface {
	int16 | uint16
	// String()
}

func Add[T Sixteen](a, b T) T {
	return a + b
}

// ./scratch.go:15:15: int does not implement Sixteen (missing String method)
func main() {
	result := Add(int16(1), int16(2))
	log.Println(result)
}
