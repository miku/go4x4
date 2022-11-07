package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// 1. Create map mapping strings to int (i.e. a counter)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		// 2. Split line into tokens, by splitting on a single whitespace (see: https://pkg.go.dev/strings#Split)

		// 3. For each token, increment its count in the map.

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for k, v := range m {
		fmt.Printf("%d\t%v\n", v, k)
	}
}
