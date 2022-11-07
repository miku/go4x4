package main

import "log"

func main() {
	var v comparable // cannot use type comparable outside a type constraint
	log.Println(v)

	// https://github.com/golang/go/issues/51338, but declined:
	// https://github.com/golang/go/issues/49587#issuecomment-995072881
}
