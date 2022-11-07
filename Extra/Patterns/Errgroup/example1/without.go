package main

import (
	"context"
	"runtime"
)

func main() {
	N := runtime.NumCPU()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pageCh := make(chan Page, N)
	doneCh := make(chan struct{})
	defer close(doneCh)
	errCh := make(chan error, 1)
	defer close(errCh)

	for i := 0; i < N; i++ {
		go g.renderPage(ctx, library, pageCh, errCh, doneCh)
	}

	go func() {
		defer close(pageCh)
		for _, page := range library.Pages {
			pageCh <- page
		}
	}()

	for N > 0 {
		select {
		case _, ok := <-doneCh:
			if ok {
				N--
			}
		case err := <-errCh:
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return nil
}
