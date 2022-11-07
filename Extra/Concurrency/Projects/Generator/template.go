package main

import (
	"fmt"
	"math/rand"
)

func randomString() string {
	return fmt.Sprintf("s-%d", rand.Intn(999999))
}

func generateIdentifier() chan string {
	ch := make(chan string)
	go func() {
		i := 0
		for {
			id := fmt.Sprintf("%d-%s", i, randomString())
			ch <- id
			i++
		}
	}()
	return ch
}

func main() {
	for v := range generateIdentifier() {
		fmt.Println(v)
	}
}
