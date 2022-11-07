package main

import "fmt"

// Write a function that prints out an ASCII table. Make use of the verbs for
// decimal, hexadecimal and octal integer representations.
// https://pkg.go.dev/fmt#hdr-Printing
//
// Output should similar to this:
//
// $ go run tasks/asciitable/main.go | column -c 60
// 32  20  40      64  40  100 @   96  60  140 `
// 33  21  41 !    65  41  101 A   97  61  141 a
// 34  22  42 "    66  42  102 B   98  62  142 b
// 35  23  43 #    67  43  103 C   99  63  143 c
// 36  24  44 $    68  44  104 D   100  64  144 d
// 37  25  45 %    69  45  105 E   101  65  145 e
// 38  26  46 &    70  46  106 F   102  66  146 f
// 39  27  47 '    71  47  107 G   103  67  147 g
// 40  28  50 (    72  48  110 H   104  68  150 h
// 41  29  51 )    73  49  111 I   105  69  151 i

func main() {
	for i := 32; i < 127; i++ {
		fmt.Printf("% 3d % 3x % 3o %c\n", i, i, i, i)
	}
}
