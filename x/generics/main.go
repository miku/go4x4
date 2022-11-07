package main

import "log"

// ./main.go:5:11: undeclared name T for array length
type Set [T]struct {
}

func main() {
	var s Set
	log.Println(s)
}
