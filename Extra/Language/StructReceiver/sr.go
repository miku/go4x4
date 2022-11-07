package sr

import (
	"math"
	"time"
)

func f(i int) int {
	return i + 1
}

func g(i float64) float64 {
	return i / 2.345
}

func h(i float64) float64 {
	time.Sleep(1 * time.Microsecond)
	return (i * i * 0.1234) / (2.345 * math.Sqrt(i))
}
