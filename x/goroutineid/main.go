package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {
	go func() {
		fmt.Println(getGID())
		time.Sleep(100 * time.Millisecond)
	}()
	time.Sleep(500 * time.Second)
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
