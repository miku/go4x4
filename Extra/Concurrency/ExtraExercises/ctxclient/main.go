package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Create a new context
	// With a deadline of 100 milliseconds
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	// Make a request, that will call the google homepage
	req, _ := http.NewRequest(http.MethodGet, "http://google.com", nil) // http.NewRequestWithContext, 1.13
	// Associate the cancellable context we just created to the request
	req = req.WithContext(ctx) // 1.7

	// Create a new HTTP client and execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	// Print the statuscode if the request succeeds
	fmt.Println("Response received, status code:", resp.StatusCode)
}
