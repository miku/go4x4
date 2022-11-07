package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
)

type reader struct{}

func (r *reader) Read(p []byte) (int, error) {
	if rand.Float64() < 0.1 {
		return 0, io.EOF
	}
	for i := 0; i < len(p); i++ {
		p[i] = '.'
	}
	return len(p), nil
}

func main() {
	rand.Seed(1)
	if _, err := io.Copy(os.Stdout, &reader{}); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}
