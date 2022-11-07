package main

import (
	"fmt"
	"log"
)

type Foo interface {
	Bar(x int) int
}

type Base struct{}

func (*Base) Bar(x int) int {
	return x
}

func Run(f Foo) {
	fmt.Println(f.Bar(99))
}

func main() {
	log.Println("example software under test")
}
