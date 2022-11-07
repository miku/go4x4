package main

import (
	"context"
	"log"
)

func Query(ctx context.Context) (string, error) {
	log.Println("Query", ctx.Value("somekey"))
	return "", nil
}

func main() {
	ctx := context.WithValue(context.Background(), "somekey", "somevalue")
	result, err := Query(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("main", ctx.Value("somekey"))
	log.Println(result)
}
