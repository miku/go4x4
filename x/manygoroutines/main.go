package main

import (
	"flag"
	"log"
	"runtime"
	"time"
)

var (
	N     = flag.Int("n", 10, "number of goroutines to start")
	sleep = flag.Duration("s", 100*time.Millisecond, "how long each goroutine sleeps")
)

func main() {
	flag.Parse()
	log.Printf("starting %d goroutines", *N)
	started := time.Now()
	for i := 0; i < *N; i++ {
		go f()
	}
	log.Printf("%d/%d goroutines started/running in %v", *N, runtime.NumGoroutine(), time.Since(started))
}

func f() {
	time.Sleep(*sleep)
}
