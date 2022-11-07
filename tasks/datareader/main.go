package main

import (
	"math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// Implement an interface type and two concrete types that implement that
// interface.

// 1. Implement an interface "DataReader" that has a single method with
// signature "Read() string"

// 2. Implement a "StringReader" (struct) that satisfies the "DataReader"
// interface and prints out a string, that the user can choose, e.g. usage would
// be like: `r := NewStringReader("hello")`

// 3. Implement a "RandomReader" that satisfies the "DataReader" interface as
// well. It should return a random string of a fixed length, e.g. length 8.

func main() {
	// 4. Declare a variable of interface type "DataReader" and assign e.g. a
	// "RandomReader" to it.

	// 5. Call "Read" on the type and print out the value.

	// 6. Use a "type switch" to check which kind of reader was used and print
	// out a short log message describing the type ("random" or "string" is enough).

}
