package main

import "fmt"

func main() {
	// (type, value)

	var x *ErrorX = nil

	var err error = x

	if err == nil {
		fmt.Println("nope, this won't happen ")
	} else {
		var reallyNil error
		fmt.Printf("here's my non-nil error: %#v %T\n", x, x)
		fmt.Printf("really nil: %T %v\n", reallyNil, reallyNil)
	}
}

type ErrorX struct{}

// "Error" interface
func (x *ErrorX) Error() string {
	return "hi i'm ErrorX"
}

func (x *ErrorX) Hello() string {
	return "Hello"
}
