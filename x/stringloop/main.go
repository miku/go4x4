package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println(sample)

	// UTF-8 is an efficient encoding for Unicode, we use only 8 bytes for 6 runes.
	fmt.Printf("len=%d, runecount=%d\n", len(sample), utf8.RuneCountInString(sample))

	// We can iterate over each byte.
	fmt.Printf("index: ")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%c [%T]", sample[i], sample[i])
	}
	// We can iterate over the runes.
	fmt.Printf("\nrange: ")
	for _, c := range sample {
		fmt.Printf("%c [%T]", c, c)
	}
	fmt.Println()
	// $ go run stringloop/main.go
	// 	= ⌘
	// len=8, runecount=6
	// index: ½ [uint8]² [uint8]= [uint8]¼ [uint8]  [uint8]â [uint8] [uint8]
	// range: � [int32]� [int32]= [int32]� [int32]  [int32]⌘ [int32]
}
