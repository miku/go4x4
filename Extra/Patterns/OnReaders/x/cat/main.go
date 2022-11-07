// A noop command or filter, that passes input from stdin to stdout.
//
// XXX: Command terminated by signal 13
//
// Demostrates the effect of buffer sizes in copy.
//
// $ make runb SIZE=10000000
// cat /dev/zero | head -c 10000000 | ./noopfilter -b 1 > /dev/null
// 2019/10/12 13:16:17 10000000 bytes filtered with buffer 1 in 14.608280051s (0.65 MB/s)
//
// $ make runb B=1073741824 SIZE=107374182400
// cat /dev/zero | head -c 107374182400 | ./noopfilter -b 1073741824 > /dev/null
// 2019/10/12 13:22:44 107374182400 bytes filtered with buffer 1073741824 in 1m27.639027456s (1168.43 MB/s)
//
// $ make runb B=512 SIZE=1073741824
// cat /dev/zero | head -c 1073741824 | ./noopfilter -b 512 > /dev/null
// 2019/10/12 13:23:47 1073741824 bytes filtered with buffer 512 in 4.014811105s (255.06 MB/s)
//
// XXX: Plot the performance per buffer size.
//
// The head command can be a bottleneck. The memory access.
//
// Reading a file directory into noopfilter, we get something like 1.5GB/s throughput.
//
// $ cat largefile | ./noopfilter -b 1073741824 > /dev/null
// 2019/10/12 13:31:23 10737418240 bytes filtered with buffer 1073741824 in 6.548966618s (1563.61 MB/s)
package main

import (
	"flag"
	"io"
	"log"
	"os"
	"time"
)

var bufSize = flag.Int("b", 0, "use specific buffer size for CopyBuffer")

// rate returns average per MB/second.
func rate(read int64, t time.Duration) float64 {
	return float64(read) / (1024 * 1024) / t.Seconds()
}

func main() {
	flag.Parse()
	if *bufSize > 0 {
		buf := make([]byte, *bufSize)
		start := time.Now()
		if n, err := io.CopyBuffer(os.Stdout, os.Stdin, buf); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("%d bytes filtered with buffer %d in %s (%0.2f MB/s)", n, *bufSize, time.Since(start), rate(n, time.Since(start)))
		}
	} else {
		start := time.Now()
		if n, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("%d bytes filtered in %s (%0.2f MB/s)", n, time.Since(start), rate(n, time.Since(start)))
		}
	}
}
