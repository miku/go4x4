// Program to feed URLs and display status code, runs parallel requests.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type workResult struct {
	Link       string `json:"link"`
	StatusCode int    `json:"status"`
	Err        error  `json:"err"`
}

var client = http.Client{
	Timeout: 5 * time.Second,
}

func worker(queue chan string, out chan workResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for line := range queue {
		resp, err := client.Get(line)
		wr := workResult{Link: line, Err: err}
		if err == nil {
			wr.StatusCode = resp.StatusCode
		}
		out <- wr
	}
}

func writer(out chan workResult, done chan bool) {
	defer func() {
		done <- true
	}()
	for wr := range out {
		b, _ := json.Marshal(wr)
		fmt.Println(string(b))
	}
}

func main() {

	var (
		numWorkers = 200
		br         = bufio.NewReader(os.Stdin)
		queue      = make(chan string)
		out        = make(chan workResult)
		done       = make(chan bool)
		wg         sync.WaitGroup
	)

	for i := 0; i < numWorkers; i++ {
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
