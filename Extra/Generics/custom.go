package main

import "log"

type Greeter interface {
	Greet() string
}

func Hello[G Greeter](g G) string {
	return g.Greet()
}

type MyGreeter struct{}

func (h *MyGreeter) Greet() string {
	return "hello"
}

func main() {
	log.Println(Hello(&MyGreeter{}))
}
