package main

import (
	"fmt"
	"log"
	"strings"
)

func isEven(x int) bool {
	return x%2 == 0
}

func divmod(a, b int) (div int, rem int) {
	div = a / b
	rem = a % b
	return div, rem
}

type MyString string

type stringFunc func(string) string

func PrefixedString(s string) string {
	return fmt.Sprintf("P" + s)
}

func main() {
	log.Println(divmod(12, 4))
	var f stringFunc
	f = strings.ToUpper
	log.Printf("%T", f)

	f = func(s string) string {
		return "anything goes"
	}
	log.Println(f("abc"))
}
