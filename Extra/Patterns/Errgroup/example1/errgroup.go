package main

import (
	"runtime"

	"golang.org/x/sync/errgroup"
)

func main() {
	N := runtime.NumCPU()
	pageCh := make(chan Page, N)
	go func() {
		defer close(pageCh)
		for _, page := range library.Pages {
			pageCh <- page
		}
	}()

	eg, ctx := errgroup.WithContext(ctx)
	for i := 0; i < N; i++ {
		eg.Go(func() error { return g.renderPage(ctx, library, pageCh) })
	}

	return eg.Wait()
}
