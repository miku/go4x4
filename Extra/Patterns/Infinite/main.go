package main

import (
	"io"
	"log"
	"os"
)

type MyReader struct{}

func (r MyReader) Read(p []byte) (int, error) {
	for i := 0; i < len(p); i++ {
		p[i] = '\n'
	}
	return len(p), io.EOF
}

func main() {
	if _, err := io.Copy(os.Stdout, MyReader{}); err != nil {
		log.Fatal(err)
	}
}
