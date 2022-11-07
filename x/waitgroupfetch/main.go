package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.spartakiade.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			resp, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			fmt.Printf("[%d] %s\n", resp.StatusCode, url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}
