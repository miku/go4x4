package main

import "log"

func Run[T []any](vs T) {
	log.Println(vs)
}

func main() {
	// []string does not implement []any ([]string missing in []any)
	vs := []string{"a", "b", "c"}
	Run(vs)
}
