package main

import (
	"io"
	"log"
	"strings"
)

type devNull struct{}

func (w *devNull) Write(p []byte) (int, error) {
	return len(p), nil
}

func main() {
	if n, err := io.Copy(&devNull{}, strings.NewReader("Hello World")); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("%d bytes copied", n)
	}
}
