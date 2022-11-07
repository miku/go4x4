package main

import "log"

// Let's write a program that takes a list of URLs and retrieves them in
// parallel.
//
// We can use http.Get to retrieve a URL:
//
//     resp, err := http.Get("http://heise.de")
//     if err != nil {
//
//     }
//     defer resp.Body.Close()
//     statusCode := resp.StatusCode
//
// --------

// Let's think about the different stages.

// 1. We could start as many goroutines as we have links, but that's not
// bounded. So let's try a worker pool.

// 2. Start N goroutines that will be our workers. There workers will fetch the
// URL.

// 3. Each worker will need to receive "work", the sender wants to signal when
// work is done.

// 4. Further options: We could use a fan-out/fan-in approach - that is we
// distribute the work across many workers, but collect the results in a single
// place, e.g. to summarize the result or send it off in a batch.

func main() {
	urls := []string{
		"http://www.youtube.com",
		"http://www.facebook.com",
		"http://www.baidu.com",
		"http://www.yahoo.com",
		"http://www.amazon.com",
		"http://www.wikipedia.org",
		"http://www.qq.com",
		"http://www.google.co.in",
		"http://www.twitter.com",
		"http://www.live.com",
	}
	log.Printf("%d urls", len(urls))
}
