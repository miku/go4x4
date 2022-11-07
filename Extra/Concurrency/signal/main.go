package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {

	c := make(chan os.Signal, 1)
	done := make(chan bool)

	signal.Notify(c, os.Interrupt)

	go func() {
		sig := <-c
		signal.Stop(c)
		fmt.Println("received signal", sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exit")
}
