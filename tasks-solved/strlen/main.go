package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var s = "Äther"

	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString((s)))

	fmt.Println()
	for i, c := range s {
		fmt.Printf("%d %c\n", i, c)

	}
}

// 6 5
//
// For strings, the range does more work for you, breaking out individual
// Unicode code points by parsing the UTF-8. Erroneous encodings consume one
// byte and produce the replacement run
//
// 0 Ä
// 2 t
// 3 h
// 4 e
// 5 r
