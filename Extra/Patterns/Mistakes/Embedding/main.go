package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Event struct {
	ID int
	time.Time
}

func main() {
	event := Event{ID: 10, Time: time.Now()}
	b, err := json.Marshal(event)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b)) // "2021-09-03T07:29:55.091376348+02:00"
}
