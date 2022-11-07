package main

import (
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
)

// XXX: Why does this deadlock?

func main() {
	pr, pw := io.Pipe()

	// A group of one?
	var wg sync.WaitGroup
	wg.Add(1)

	log.Printf("G1=%d", runtime.NumGoroutine())

	go func() {
		defer wg.Done()
		if _, err := io.Copy(os.Stdout, pr); err != nil {
			log.Fatal(err)
		}
		// There seems to be a scheduling point here.
		log.Println("Exiting")
	}()

	log.Printf("G2=%d", runtime.NumGoroutine())

	go func() {
		sr := strings.NewReader("Hello World")
		if _, err := io.Copy(pw, sr); err != nil {
			log.Fatal(err)
		}
	}()

	// We cannot guaratee, that this is executed after wg.Done(), so we get a
	// deadlock. This waits already, but the other thread hat not been
	// scheduled yet.

	log.Printf("G3=%d", runtime.NumGoroutine())

	// We cannot thread this, either. We need to thread out the write.
	wg.Wait()

	// What does wg.Wait do to the scheduler? You could think it would stop this G and let the other run?
}
