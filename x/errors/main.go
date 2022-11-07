package main

import (
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := io.Copy(os.Stdout, f); err != nil {
		log.Fatal(err)
	}
}
