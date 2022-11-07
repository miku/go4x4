// Cancellation example without context.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

type Task struct {
	closed chan struct{}
	wg     sync.WaitGroup
	ticker *time.Ticker
}

func (t *Task) Run() {
	for {
		select {
		case <-t.closed:
			return
		case <-t.ticker.C:
			handle()
		}
	}
}

func (t *Task) Stop() {
	close(t.closed)
	t.wg.Wait()
}

func handle() {
	for i := 0; i < 5; i++ {
		fmt.Print("#")
		time.Sleep(time.Millisecond * 200)
	}
	fmt.Println()
}

func main() {
	task := &Task{
		closed: make(chan struct{}),
		ticker: time.NewTicker(time.Second * 2),
	}
	task.wg.Add(1)
	go func() {
		defer task.wg.Done()
		task.Run()
	}()
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	select {
	case sig := <-c:
		singal.Stop(c)
		fmt.Printf("Got %s signal. Aborting...\n", sig)
		task.Stop()
	}
}
