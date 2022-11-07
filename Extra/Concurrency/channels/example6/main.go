// Parallel link checker.
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
)

type result struct {
	Link       string `json:"link,omitempty"`
	StatusCode int    `json:"code,omitempty"`
	Error      error  `json:"error,omitempty"`
}

func (r result) WriteTo(w io.Writer) (int64, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return 0, err
	}
	b = append(b, '\n')
	n, err := w.Write(b)
	return int64(n), err
}

func worker(work chan string, out chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for link := range work {
		var buf bytes.Buffer
		resp, err := http.Get(link)
		if err != nil {
			r := result{Link: link, Error: err}
			if _, err := r.WriteTo(&buf); err != nil {
				log.Fatal(err)
			}
			continue
		}
		defer resp.Body.Close()
		r := result{Link: link, StatusCode: resp.StatusCode}
		if _, err := r.WriteTo(&buf); err != nil {
			log.Fatal(err)
		}
		out <- buf.String()
	}
}

func writer(out chan string, done chan bool) {
	for line := range out {
		fmt.Printf(line)
	}
	done <- true
}

func main() {
	f, err := os.Open("small.txt")
	if err != nil {
		log.Fatal(err)
	}
	br := bufio.NewReader(f)

	queue := make(chan string)
	out := make(chan string)
	done := make(chan bool)

	var wg sync.WaitGroup

	// Start some workers.
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go worker(queue, out, &wg)
	}

	log.Printf("NumGoroutines=%v", runtime.NumGoroutine())

	// Start the fan-in go routine.
	go writer(out, done)

	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		link := strings.TrimSpace(line)
		queue <- link
	}

	close(queue)
	wg.Wait()
	close(out)
	<-done
}
