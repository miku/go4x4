package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Entry struct {
	Name string    `json:"name"`
	Date time.Time `json:"t"`
	Refs []string  `json:"refs"`
}

func main() {
	entry := Entry{
		Name: "Go",
		Date: time.Now(),
		Refs: []string{"a", "b", "c"},
	}
	b, err := json.Marshal(entry)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

// {
// 	"name": "Go",
// 	"t": "2022-05-03T20:56:11.062689686+02:00",
// 	"refs": [
// 	  "a",
// 	  "b",
// 	  "c"
// 	]
//   }
