package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(len("Ä"))                    // 2
	fmt.Println(utf8.RuneCountInString("Ä")) // 1
	fmt.Println(int('Ä'))                    // 196
	fmt.Println('Ä' == 196)
	fmt.Println(int('日')) // 26085

	for i, c := range "Ä" {
		fmt.Printf("%d c=%c T=%T v=%v d=%d x=%x\n", i, c, c, c, c, c)
	}
}

// 2
// 1
// 196
// true
// 26085
// 0 c=Ä T=int32 v=196 d=196 x=c4
