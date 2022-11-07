package main

import "fmt"

type Greeter interface {
	Greet() string
}

type A struct{}

func (a *A) Greet() string {
	return "Hey from A"
}

type B struct{}

func (b B) Greet() string {
	return "Hello from B"
}

func main() {
	// var g Greeter = A{}
	// var g Greeter = &A{}
	// var g Greeter = B{}
	// var g Greeter = &B{}

	fmt.Println(g.Greet())
}
