// All material is licensed under the GNU Free Documentation License
// https://github.com/gobridge/gotraining/blob/master/LICENSE

// http://play.golang.org/p/sKLLsUa5hH

// export GODEBUG=schedtrace=1000
// export GODEBUG=schedtrace=1000,scheddetail=1

// Sample program to review scheduler stats.
//
// $ make
// $ GODEBUG=schedtrace=1000 ./main
package main

import (
	"sync"
	"time"
)

// Create a waitgroup.
var wg sync.WaitGroup

// main is the entry point for the application.
func main() {
	N := 20
	// We are going to create 20 goroutines.
	wg.Add(N)

	// Create those 10 goroutines.
	for i := 0; i < N; i++ {
		go f()
	}

	// Wait for all the goroutines to complete.
	wg.Wait()

	// Wait to see the global runqueue deplete.
	time.Sleep(3 * time.Second)
}

// f does some CPU bound work.
func f() {
	time.Sleep(time.Second)

	var count int
	for i := 0; i < 1e10; i++ {
		count++
	}

	wg.Done()
}
