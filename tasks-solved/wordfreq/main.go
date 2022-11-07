package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// 1. Create map mapping strings to int (i.e. a counter)
	m := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// 2. Split line into tokens
		tokens := strings.Split(line, " ")
		// 3. Increment count for each token.
		for _, t := range tokens {
			t = strings.TrimSpace(t)
			if len(t) == 0 {
				continue
			}
			m[t]++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for k, v := range m {
		fmt.Printf("%d\t%v\n", v, k)
	}
}
