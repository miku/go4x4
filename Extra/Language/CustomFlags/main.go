package main

import (
	"flag"
	"fmt"
	"strings"
)

// ArrayFlags allows to store lists of flag values.
type Array []string

// String representation.
func (f *Array) String() string {
	return strings.Join(*f, ", ")
}

// Set appends a value.
func (f *Array) Set(value string) error {
	*f = append(*f, value)
	return nil
}

func main() {
	var names Array
	flag.Var(&names, "n", "name, repeatable")
	flag.Parse()
	for i, name := range names {
		fmt.Printf("%d %s\n", i, name)
	}
}
