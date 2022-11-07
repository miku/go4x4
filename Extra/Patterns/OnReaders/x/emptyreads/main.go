// How many empty reads will io.Copy allow?
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

var counter int

type ReaderFunc func([]byte) (int, error)

func (r ReaderFunc) Read(p []byte) (int, error) {
	return r(p)
}

// emptyRead returns 0 and nil.
func emptyRead(p []byte) (int, error) {
	counter++
	return 0, nil
}

// MustCopy allows only a fixed number of empty reads (e.g. like scanner)?
func MustCopy(w io.Writer, r io.Reader) (int, error) {
	buf := make([]byte, 4096)
	n, err := r.Read(buf)
	if err == nil && n == 0 {
		return 0, fmt.Errorf("empty read")
	}
	return w.Write(buf)
}

// MustCopy allows only a fixed number of empty reads (e.g. like scanner)?
func MustCopyRetry(w io.Writer, r io.Reader, maxRetries int) (int, error) {
	var k int
	buf := make([]byte, 4096)
	for {
		n, err := r.Read(buf)
		if err == nil && n == 0 {
			if k < maxRetries {
				continue
			} else {
				return 0, fmt.Errorf("empty read")
			}
		}
		break
	}
	return w.Write(buf)
}

func main() {
	defer func() {
		log.Printf("%d reads", counter)
	}()
	if _, err := MustCopy(ioutil.Discard, ReaderFunc(emptyRead)); err != nil {
		log.Fatal(err)
	}
}
