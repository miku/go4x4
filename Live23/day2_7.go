package main

import "log"

type Greeter interface {
	Greet() string
}

type Hello struct{}

func (h Hello) Greet() string {
	return "Hello!"
}

type Hi struct{}

func (h Hi) Greet() string {
	return "Hi!"
}

func main() {
	var g Greeter

	hello := Hello{}
	hi := Hi{}
	log.Printf("%T", hello)
	log.Printf("%T", hi)

	g = hello
	log.Printf("%v", g.Greet())
	g = hi
	log.Printf("%v", g.Greet())
}
