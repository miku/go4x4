package main

import (
	"log"
	"time"
)

func main() {
	var (
		a = "hello"
		t = time.Now().Second()
	)
	switch {
	case len(a) > 10:
		log.Println("[1]")
	case t%2 == 0 && a == "hello":
		log.Println("[2]")
	default:
		log.Println("[default]")
	}
}
