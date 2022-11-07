package main

import (
	"io"
	"sync"
)

type File struct {
	sync.Mutex
	rw io.ReadWriter
}

func main() {
	f := File{}
	f.Lock()
}
