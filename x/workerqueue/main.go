package main

import (
	"fmt"
)

func worker(id int, queue chan string) {
	for w := range queue {
		fmt.Printf("%d %s\n", id, w)
	}
	fmt.Printf("worker %d done\n", id)
}

func main() {
	ch := make(chan string)

	for i := 0; i < 3; i++ {
		go worker(i, ch)
	}

	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("item %d", i)
	}
	close(ch)
	fmt.Println("press enter to exit")
	fmt.Scanln() // wait for Enter Key
}
