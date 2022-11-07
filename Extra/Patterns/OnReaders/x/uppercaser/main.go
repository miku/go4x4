package main

import (
	"bytes"
	"io"
	"log"
	"os"
)

type UpperReader struct {
	r io.Reader
}

func (r *UpperReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	copy(p, bytes.ToUpper(p)) // Breaks on unicode.
	return n, err
}

func main() {
	if n, err := io.Copy(os.Stdout, &UpperReader{os.Stdin}); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("copied %d bytes", n)
	}
}
