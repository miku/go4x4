package main

import "sync"

type Task struct {
	sync.Mutex
	state int
}

func main() {
	task := Task{}
	task.Lock()
	task.state = 1
	task.Unlock()
}
