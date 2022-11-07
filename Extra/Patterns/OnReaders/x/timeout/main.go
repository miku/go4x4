package main

import (
	"errors"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// ErrTimeout signals a timeout.
var ErrTimeout = errors.New("timeout")

// TimeoutReader times out, if read takes too long.
// https://github.com/golang/go/wiki/Timeouts
type TimeoutReader struct {
	r       io.Reader
	timeout time.Duration
}

type readResult struct {
	b   []byte
	err error
}

func (r *TimeoutReader) Read(p []byte) (n int, err error) {
	ch := make(chan readResult, 1)

	go func() {
		pp := make([]byte, len(p))
		_, err := r.r.Read(pp)
		ch <- readResult{pp, err}
	}()

	select {
	case <-time.After(r.timeout):
		return 0, ErrTimeout
	case res := <-ch:
		copy(p, res.b)
		return len(p), res.err
	}
}

// Slow is a sleepy reader. Used for testing.
type Slow struct {
	Sleep time.Duration
}

// Read returns a few bytes after a given duration.
func (r *Slow) Read(p []byte) (n int, err error) {
	time.Sleep(r.Sleep)
	for i := 0; i < len(p); i++ {
		p[i] = '\x7a'
	}
	return len(p), io.EOF
}

func main() {
	var r io.Reader

	// Strings are fast.
	r = &TimeoutReader{
		r:       strings.NewReader("Hello World!\n"),
		timeout: 100 * time.Millisecond,
	}
	io.Copy(os.Stdout, r)

	// Slow reader will timeout, if Slow sleeps to long.
	r = &TimeoutReader{
		r:       &Slow{Sleep: 500 * time.Millisecond},
		timeout: 100 * time.Millisecond,
	}
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
	// Hello World!
	// 2017/01/19 19:56:29 timeout
	// exit status 1
}
