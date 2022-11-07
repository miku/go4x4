package main

import (
	"context"
	"log"
	"time"
)

func Query(ctx context.Context) (string, error) {
	// this select block until one case is available
	var i int
	for {
		i++
		select {
		case <-time.After(200 * time.Millisecond):
			return "ok", nil
		case <-ctx.Done(): // read from a closed channel returns the zero value immediatly
			log.Println("cancel requested ...")
			// time.Sleep(200 * time.Millisecond)
			return "", ctx.Err()
		}
	}
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 500*time.Millisecond)
	// defer cancel()
	// Query would normally take a second, but we're cancelling it.
	result, err := Query(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)

	// result, err = Query(ctx)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(result)
	//
	// result, err = Query(ctx) // "read from close channel"
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(result)
}
