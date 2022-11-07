package main

import (
	"fmt"
	"unicode"
)

func main() {
	chars := []rune{'a', 'α', '\u29c9'}
	for _, c := range chars {
		fmt.Printf("%c %v\n", c, unicode.IsLetter(c))
	}
	// a true
	// α true
	// ⧉ false
}
