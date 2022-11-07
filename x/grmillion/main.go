package main

import (
	"log"
	"runtime"
	"time"
)

func main() {
	log.Println("starting goroutines ...")
	for i := 0; i < 1_000_000; i++ {
		go func() {
			time.Sleep(10 * time.Second)
		}()
	}
	log.Printf("waiting before shutdown (running=%d)", runtime.NumGoroutine())
	time.Sleep(5 * time.Second)
	log.Println("done")
	// $ go run x/grmillion/main.go
	// 2022/05/03 15:35:31 starting goroutines ...
	// 2022/05/03 15:35:33 waiting before shutdown (running=1000001)
	// 2022/05/03 15:35:38 done
}
