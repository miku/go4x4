package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func RandomString(length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteRune(rune(97 + rand.Intn(26)))
	}
	// log.Println(sb.String())
	return sb.String()
}
func generateId() <-chan string {
	ch := make(chan string)
	go func() {
		for {
			s := RandomString(11)
			ch <- s
		}
	}()
	return ch
}

func main() {

	for v := range generateId() {
		fmt.Println(v)
	}
}
