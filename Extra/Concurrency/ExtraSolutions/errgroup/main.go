package main

import (
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
			// return fmt.Errorf("failed to fetch url")
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		log.Println("Successfully fetched all URLs.")
	} else {
		log.Printf("failed with: %v", err)
	}
}
