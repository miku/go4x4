package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var a []*string
	for i := 0; i < 10; i++ {
		s := strconv.Itoa(i)
		a = append(a, &s)
	}
	i, j := 2, 7
	dump(a) // [0123456789]

	copy(a[i:], a[j:]) // copies 789 after 01, result in [01789...]

	// explicitly set remaining values to nil
	for k, n := len(a)-j+i, len(a); k < n; k++ {
		log.Printf("setting s[%d] to nil", k)
		a[k] = nil // or the zero value of T
	}
	// shorten the slice
	a = a[:len(a)-j+i]
	dump(a)

	// no values remain, hence a "safe cut"
	a = a[:10]
	dump(a)
}

func dump(ss []*string) {
	fmt.Printf("[")
	for _, s := range ss {
		if s == nil {
			fmt.Printf("N")
		} else {
			fmt.Printf("%s", *s)
		}
	}
	fmt.Printf("]")
	fmt.Println()
}
