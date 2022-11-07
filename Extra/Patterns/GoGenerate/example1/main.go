package main

import "log"

//go:generate bash -c "ls -lah > sample.txt"

func main() {
	log.Println("hello world")
}
