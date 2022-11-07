package main

import (
	"fmt"
	"runtime"
	"time"
)

func f() {
	fmt.Println("hello world")
}

func main() {
	fmt.Println(runtime.NumGoroutine())
	go f()
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(1 * time.Second)
}
