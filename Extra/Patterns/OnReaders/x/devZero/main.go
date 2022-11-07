package main

import (
	"io"
	"log"
	"os"
)

type devZero struct{}

func (r *devZero) Read(p []byte) (int, error) {
	for i := 0; i < len(p); i++ {
		p[i] = '\x00'
	}
	return len(p), nil
}

func main() {
	if _, err := io.CopyN(os.Stdout, &devZero{}, 10); err != nil {
		log.Fatal(err)
	}
}
