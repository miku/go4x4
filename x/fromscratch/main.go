package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	flag.Parse()
	url := "https://heise.de"
	if flag.NArg() > 0 {
		url = flag.Arg(0)
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatal(err)
	}
}
