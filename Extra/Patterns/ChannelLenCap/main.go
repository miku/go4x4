package main

import "log"

func main() {
	ch1 := make(chan int)
	log.Println("ch1", len(ch1), cap(ch1))

	ch2 := make(chan int, 10)
	log.Println("ch2", len(ch2), cap(ch2))

	ch2 <- 1
	log.Println("ch2", len(ch2), cap(ch2))
}
