// $ go run main.go | wc -l
// 9252
//
// More proof that this works: trickle (a lightweight userspace bandwidth shaper)
//
// $ trickle -d 10 wget https://dl.google.com/go/go1.18.windows-amd64.zip # 4h+
// $ time trickle -d 10 go run main.go | wc -l
// 13131
//
// real    0m1.831s
// user    0m1.046s
// sys     0m0.351s
//
package main

import (
	"archive/zip"
	"fmt"
	"log"
	"net/http"

	bufra "github.com/avvmoto/buf-readerat"
	"github.com/snabb/httpreaderat"
)

func main() {
	req, err := http.NewRequest("GET", "https://dl.google.com/go/go1.18.windows-amd64.zip", nil)
	// To see the transmitted header, uncomment:
	// req, err := http.NewRequest("GET", "https://webhook.site/8b236e4f-b35c-49ac-95ac-29c3367be2fa", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Package httpreaderat implements io.ReaderAt that makes HTTP Range Requests.
	// It can be used for example with "archive/zip" package in Go standard
	// library. Together they can be used to access remote (HTTP accessible)
	// ZIP archives without needing to download the whole archive file.
	htrdr, err := httpreaderat.New(nil, req, nil) // client *http.Client, req *http.Request, bs Store
	if err != nil {
		log.Fatal(err)
	}

	// Package buf-readerat implements buffered io.ReaderAt. It wraps an io.ReaderAt
	// object, creating another io.ReaderAt object that also implements the interface
	// but provides buffering.
	bhtrdr := bufra.NewBufReaderAt(htrdr, 1024*1024)

	// https://golang.org/pkg/archive/zip/#NewReader
	zrdr, err := zip.NewReader(bhtrdr, htrdr.Size())
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range zrdr.File {
		fmt.Println(f.Name)
	}
}
