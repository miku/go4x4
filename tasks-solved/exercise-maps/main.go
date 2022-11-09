package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, w := range strings.Split(s, " ") {
		m[w]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
