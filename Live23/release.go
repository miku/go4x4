package main

import (
	"encoding/json"
	"log"
	"os"
)

type Release struct {
	ExtIds struct {
		Doi string `json:"doi"`
	} `json:"ext_ids"`
	WorkId string `json:"work_id"`
}

func main() {
	var r Release
	_ = json.NewDecoder(os.Stdin).Decode(&r)
	log.Println(r)
}
