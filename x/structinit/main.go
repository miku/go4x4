package main

import "fmt"

type Location struct {
	Lat  float64
	Long float64
}

type Peer struct {
	Name     string
	Location Location
}

func main() {
	peer := Peer{
		Name: "martin",
		Location: Location{
			Lat:  51.33962,
			Long: 12.37129,
		}}
	fmt.Println(peer)
	// {martin {51.33962 12.37129}}
}
