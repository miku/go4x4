package main

import (
	"sync"
)

func main() {

	// trace.Start(os.Stderr)
	// defer trace.Stop()

	var wg sync.WaitGroup

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			a := 0

			for i := 0; i < 1e6; i++ {
				a += 1
			}

			wg.Done()
		}()
	}

	wg.Wait()
}
