package main

import "log"

func main() {
	var i int = 8
	var j = &i // address of
	log.Println(i)
	log.Println(j)
	log.Println(*j)
	var k *int
	k = &i
	log.Println(k)
}
