package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"time"
)

// Exercise: Worker pool with fan-in and error handling. We want to pass errors
// along the results.
//
// (1) Create a result type, that hold result and error.
//
// (2) Adjust worker, fanIn function and main to use the new result type.
//
// (3) Adjust worker, so it sometimes (e.g. in 10% of the cases) will return an
// error.
//
// (4) In the fanIn function print out the good results and count the errors,
// print out the number of errors at the end.

// result value and possible error.
type result struct {
	value string
	err   error
}

// tasks returns a slice of strings simulating tasks.
func tasks() (result []string) {
	for i := 0; i < 1000; i++ {
		result = append(result, fmt.Sprintf("item-%d", i))
	}
	return
}

// worker has an id, a queue to receive tasks from, WaitGroup for joining.
func worker(id int, queue chan string, out chan result, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range queue {
		upper := fmt.Sprintf("[%d] %s", id, strings.ToUpper(task))
		var err error
		if rand.Float64() < 0.1 {
			err = fmt.Errorf("failed")
		}
		out <- result{
			value: upper,
			err:   err,
		}
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func fanIn(done chan bool, out chan result) {
	var failed int
	for result := range out {
		if result.err != nil {
			failed++
			continue
		}
		log.Printf("fanIn: %s", result.value)
	}
	log.Printf("%d tasks failed", failed)
	done <- true
}

func main() {
	queue := make(chan string)
	out := make(chan result)
	done := make(chan bool)

	var wg sync.WaitGroup

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go worker(i, queue, out, &wg)
	}

	go fanIn(done, out)

	for _, task := range tasks() {
		queue <- task
	}
	close(queue)
	wg.Wait()
	close(out)
	<-done
}
