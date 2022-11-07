package main

import (
	"encoding/json"
	"fmt"
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
	b, _ := json.Marshal(peer)
	fmt.Println(string(b))
	// {"name":"martin","loc":{"lat":51.33962,"long":12.37129}}
}
