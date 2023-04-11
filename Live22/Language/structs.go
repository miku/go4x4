package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type Peer struct {
	Name     string   `json:"name"`
	Location Location `json:"loc"`
}

func main() {
	peer := Peer{
		Name: "martin",
		Location: Location{
			Lat:  51.33962,
			Long: 12.37129,
		}}
	fmt.Println(peer)
	fmt.Println(peer.Location)
	fmt.Println(peer.Location.Long)
	// {martin {51.33962 12.37129}}
	b, err := json.Marshal(peer)
	if err != nil {
		log.Fatal(err)
	}

	b, err = json.Marshal(peer)
	if err != nil {
		log.Fatal(err)
	}

	// if err := f.Close(); err != nil {
	// }

	if b, err := json.Marshal(peer); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(b))
	}

	if 1 < 2 {
		log.Println("ok")
	}
}
