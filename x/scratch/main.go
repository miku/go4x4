package main

import (
	"errors"
	"fmt"
)

var (
	a      = "a"
	b      = "b"
	c bool = true
)

var (
	ErrNotFound = errors.New("not found")
	ErrBadData  = errors.New("bad data")
)

func main() {
	fmt.Println(a, b, c)
}
