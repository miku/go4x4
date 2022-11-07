package main

import "log"

func Print(s string) {
	log.Println(s)
}

func a() {
	status := "ok"
	defer Print(status)
	status = "fail"
}

func b() {
	status := "ok"
	defer func() {
		Print(status)
	}()
	status = "fail"
}

func main() {
	a()
	b()
}
