package main

import (
	"fmt"
	"time"
)

// Switch can be used with variables of many types, or no variable at all. Write
// a short program, that get current day and depending on the hour print "Good
// morning" (before noon), "Good afternoon" (before 5pm) or "Good evening".

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
