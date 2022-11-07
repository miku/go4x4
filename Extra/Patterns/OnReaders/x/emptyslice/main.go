package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	log.Fatal("err")
	r := strings.NewReader("Hello World")
	buf := make([]byte, 1000)
	log.Println("CopyBuffer")
	if _, err := io.CopyBuffer(os.Stdout, r, buf); err != nil {
		log.Fatal(err)
	}
}
