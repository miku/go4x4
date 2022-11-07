package main

import (
	"fmt"
)

func f() {
	x := []int{}     // literal slice, no items
	x = append(x, 0) // 1. append
	x = append(x, 1) // 2. append

	sliceInfo("x", x)

	y := append(x, 2) // 3. append, new var y
	z := append(x, 3) // 4. append, new var z

	fmt.Println("y", y) // WHAT IS Y? -- 0 1 2 | 0 1 2   | 0 1 2 3
	fmt.Println("z", z) // WHAT IS Z? -- 0 1 3 | 0 1 2 3 | 0 1 2 3

	// sliceInfo("x", x)
	// sliceInfo("y", y)
	// sliceInfo("z", z)

}

func g() {
	x := []int{}     // literal slice, no items
	x = append(x, 0) // 1. append
	x = append(x, 1) // 2. append
	x = append(x, 2) // 3. append

	sliceInfo("x", x)

	y := append(x, 3) // 4. append, new var y

	sliceInfo("x", x)
	sliceInfo("y", y)

	z := append(x, 4) // 5. append, new var z

	sliceInfo("x", x)
	sliceInfo("y", y)
	sliceInfo("z", z)

	// fmt.Println("y", y) // WHAT IS Y? | 0 1 2 3 | 0 1 2 3   | ... 000 | 0 1 2 4
	// fmt.Println("z", z) // WHAT IS Z? | 0 1 2 4 | 0 1 2 3 4 | ... 000 | 0 1 2 4
}

func main() {
	fmt.Println("f")
	f()
	// fmt.Println("g")
	// g()
}

func sliceInfo(name string, s []int) {
	fmt.Printf("[%s] %p len=%d cap=%d %v\n", name, s, len(s), cap(s), s)
}

// W3hdIDB4YzAwMDBiNjAyMCBsZW49MiBjYXA9Mgp5IFswIDEgMl0KeiBbMCAxIDNdClt4XSAweGMwMDAw
// YmUwNjAgbGVuPTMgY2FwPTQKW3hdIDB4YzAwMDBiZTA2MCBsZW49MyBjYXA9NApbeV0gMHhjMDAwMGJl
// MDYwIGxlbj00IGNhcD00Clt4XSAweGMwMDAwYmUwNjAgbGVuPTMgY2FwPTQKW3ldIDB4YzAwMDBiZTA2
// MCBsZW49NCBjYXA9NApbel0gMHhjMDAwMGJlMDYwIGxlbj00IGNhcD00CnkgWzAgMSAyIDRdCnogWzAg
// MSAyIDRdCg==
