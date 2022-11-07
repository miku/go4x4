// A teebench reads and duplicate data in memory.
//
package main

import (
	"flag"
	"io"
	"log"
	"os"
	"time"
)

var bufSize = flag.Int("b", 0, "use specific buffer size for CopyBuffer")

// counterWriter discards the data, only keeps track of bytes written. Not
// thread-safe.
type counterWriter struct {
	n int64
}

// Write discards the data, counts the bytes written.
func (w *counterWriter) Write(b []byte) (int, error) {
	w.n += int64(len(b))
	return len(b), nil
}

// rate returns average per MB/second.
func rate(read int64, t time.Duration) float64 {
	return float64(read) / (1024 * 1024) / t.Seconds()
}

func main() {
	flag.Parse()
	start := time.Now()
	cw := counterWriter{}
	if n, err := io.Copy(os.Stdout, io.TeeReader(os.Stdin, &cw)); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("%d bytes filtered with tee (%d) in %s (%0.2f MB/s)",
			n, cw.n, time.Since(start), rate(n, time.Since(start)))
	}
}
