package main

import "log"

type MyStr string

func main() {
	m := MyStr("hello")
	log.Println(m)
}
