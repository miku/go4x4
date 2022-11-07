package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	A string `json:"a"`
}

func main() {
	data := Data{}
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}
