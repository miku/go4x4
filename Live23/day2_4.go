package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
)

func isEven(x int) bool {
	return x%2 == 0
}

func divmod(a, b int) (int, int) {
	return a / b, a % b

}

func main() {

	for i := 0; i < 3; i++ {
		log.Println(i)
	}
	var j int
	for { // while True ...
		j++
		if j == 10 {
			break
		}
	}
	for j < 20 {
		j++
	}

	for {
		if j > 30 {
			break
		}
		j++
	}

	log.Printf("done, %v", j)

	if 1 > 0 {
		log.Println("1 > 0")
	}
	// short form
	x := rand.Intn(1000)
	if x > 500 {
		log.Println("x > 500")
	} else if x > 800 {
		log.Println("x > 800")
	} else {
		log.Println(x)
	}
	log.Println(x)

	x = 11
	fmt.Print("go runs on ... ")
	switch {
	case isEven(x):
		fmt.Println("isEven")
	case runtime.GOOS == "darwin":
		fmt.Println("darwin")
	default:
		fmt.Println("something else")
	}
}
