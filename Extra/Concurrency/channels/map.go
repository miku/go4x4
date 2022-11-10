package main

import (
	"log"
	"sync"
	"time"
)

type Map struct {
	m          map[int]int
	sync.Mutex // embed, "embedding a lock"
}

func main() {
	// m := make(map[int]int)

	m := Map{m: make(map[int]int)}

	for g := 0; g < 1000; g++ {
		go func(g int) {
			m.Lock()
			m.m[g] = g
			m.Unlock()
		}(g)
	}
	time.Sleep(1 * time.Second)
	log.Println(m)
}
