package main

import (
	"io"
	"log"
	"os"
)

type Endless struct{}

func (e *Endless) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		p[i] = 'z'
	}
	return len(p), nil
}

func main() {
	if _, err := io.Copy(os.Stdout, Endless{}); err != nil {
		log.Fatal(err)
	}
}
