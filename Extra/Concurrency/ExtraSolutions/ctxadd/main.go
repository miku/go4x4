package main

import (
	"context"
	"log"
	"time"
)

// Exercise: Write a function addContext, that takes a context and two integers
// and returns the result and an error. Make it so the add is articially
// delayed.
//
// (1) Write a function addContext, with a context and two integers, a, b as
// parameters. It should return the sum a + b and an error. In the function
// check for cancellation via Done, return early.
//
// (2) In main(), create a new context with a timeout, call the add function and
// display a result or an error.

func addContext(ctx context.Context, a, b int) (int, error) {
	resultC := make(chan int)

	go func() {
		time.Sleep(300 * time.Millisecond)
		resultC <- a + b
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case v := <-resultC:
		return v, nil
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()
	v, err := addContext(ctx, 2, 2)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(v)
}
