package main

import "fmt"

func main() {
	// While data is only used in loopData, it might be used elsewhere.
	data := []int{1, 2, 3, 4}

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}
