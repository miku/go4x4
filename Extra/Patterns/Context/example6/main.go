package main

import (
	"context"
	"log"
)

func g(ctx context.Context) (string, error) {
	log.Println("g", ctx.Value("main"))
	log.Println("g", ctx.Value("f"))
	return "", nil
}

func f(ctx context.Context) (string, error) {
	log.Println("f", ctx.Value("main"))
	ctx = context.WithValue(ctx, "f", "f:somevalue")
	g(ctx)
	return "", nil
}

func main() {
	ctx := context.WithValue(context.Background(), "main", "main:somevalue")
	// Query would normally take a second, but we're cancelling it.
	result, err := f(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("main", ctx.Value("main"))
	log.Println("main", ctx.Value("f")) // context is immutable
	log.Println(result)
}
