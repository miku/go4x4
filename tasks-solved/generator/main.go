package main

import "fmt"

func gen(n int) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch

}
func main() {
	for i := range gen(10) {
		fmt.Println(i)
	}
}
