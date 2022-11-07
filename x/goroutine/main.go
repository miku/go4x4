package main

import (
	"log"
	"runtime"
	"time"
)

func hello() {
	log.Println("hello")
}

func main() {
	log.Printf("cpus=%d", runtime.NumCPU())
	log.Printf("goroutines=%d", runtime.NumGoroutine())

	go hello()

	log.Printf("goroutines=%d", runtime.NumGoroutine())
	time.Sleep(10 * time.Millisecond)
}
