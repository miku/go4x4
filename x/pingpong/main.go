package main

import (
	"fmt"
	"time"
)

func player(name string, table chan int) {
	for {
		ball := <-table
		fmt.Printf("%s\n", name)
		time.Sleep(500 * time.Millisecond)
		table <- ball
	}
}

func main() {
	table := make(chan int)
	go player("ping", table)
	go player("pong", table)

	table <- 1

	time.Sleep(5 * time.Second)
}
