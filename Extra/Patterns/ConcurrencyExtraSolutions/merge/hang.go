package main

import "fmt"

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for _, v := range vs {
			c <- v
		}
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v, ok := <-a:
				if !ok {
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					continue
				}
				c <- v
			}
		}
	}()
	return c
}

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)

	for v := range merge(a, b) {
		fmt.Println(v)
	}
}
