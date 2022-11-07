package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// Exercise: A worker pool example. Some data and a basic worker is already
// there. A task is represented by a string (could be a URL for example).
//
// (1) Complete the main function, setup a number of workers, the queue.
// (2) Iterate over the data and put the strings of the queue.

// tasks returns a slice of strings simulating tasks.
func tasks() (result []string) {
	for i := 0; i < 1000; i++ {
		result = append(result, fmt.Sprintf("item-%d", i))
	}
	return
}

// worker has an id, a queue to receive tasks from, WaitGroup for joining.
func worker(id int, queue chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range queue {
		log.Printf("[%d] %s", id, strings.ToUpper(task))
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func main() {
}
