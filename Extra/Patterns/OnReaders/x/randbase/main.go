package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	enc := base64.NewEncoder(base64.RawURLEncoding, os.Stdout)
	defer enc.Close()

	// rand.Reader: On Linux and FreeBSD, Reader uses getrandom(2) if available,
	// /dev/urandom otherwise.
	if _, err := io.CopyN(enc, rand.Reader, 16); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}
