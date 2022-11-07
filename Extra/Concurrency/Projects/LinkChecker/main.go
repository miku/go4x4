package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

// workResult of a URL lookup
type workResult struct {
	Link       string `json:"link"`
	StatusCode int    `json:"status"`
	Err        error  `json:"err"`
}

var client = http.Client{
	Timeout: 5 * time.Second,
}

// worker takes work item off queue and puts result back on an out channel.
func worker(queue chan string, out chan workResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for link := range queue {
		resp, err := client.Get(link)
		r := workResult{Link: link, Err: err}
		if err == nil {
			r.StatusCode = resp.StatusCode
		}
		out <- r
	}
}

// writer takes results and writes them out, errc serves both as a way to
// singal the end (with nil) and pass the error.
func writer(out chan workResult, done chan bool) {
	for wr := range out {
		b, err := json.Marshal(wr)
		if err != nil {
			log.Printf("failed to marshal result: %v", err)
		} else {
			fmt.Println(string(b))
		}
	}
	done <- true
}

func main() {
	var (
		br    = bufio.NewReader(os.Stdin)
		queue = make(chan string)
		out   = make(chan workResult)
		done  = make(chan bool)
		wg    sync.WaitGroup
	)
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go worker(queue, out, &wg)
	}
	go writer(out, done)
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		line = strings.TrimSpace(line)
		queue <- line
	}
	close(queue)
	wg.Wait()
	close(out)
	<-done
}
