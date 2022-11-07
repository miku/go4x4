package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/errors"
)

var errReadFailed = errors.New("read failed")

func f(r io.Reader) error {
	_, err := ioutil.ReadAll(r)
	if err != nil {
		return errors.Wrap(err, "f failed")
	}
	return nil
}

type failReader struct{}

func (f *failReader) Read(p []byte) (int, error) {
	return 0, errReadFailed
}

func main() {
	if err := f(&failReader{}); err != nil {
		// we can unwrap the error (nil, when done)
		log.Printf("unwrapped: %v", errors.Unwrap(errors.Unwrap(err))) // unwrapped: read failed
		log.Println(errors.Is(err, os.ErrClosed))                      // false
		log.Println(errors.Is(err, errReadFailed))                     // true
		log.Fatalf("%v %T", err, err)                                  // f failed: read failed *errors.withStack
	}
}
