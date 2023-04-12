package main

import (
	"fmt"
	"log"
)

type Location struct {
	Lat  float64 `json:"lat"` // struct tags
	Long float64 `json:"long"`
}

type Peer struct {
	Name     string   `json:"name"`
	Location Location `json:"location"`
	active   bool
}

func (p *Peer) SetName(name string) {
	p.Name = name
}

func (p *Peer) String() string {
	return fmt.Sprintf("peer: %s", p.Name)
}

func main() {
	var loc = Location{
		Lat:  1.2,
		Long: 5.4,
	}
	var peer = &Peer{"F8", loc, true}
	log.Printf("%T", peer)
	peer.SetName("F9")
	log.Println(peer)
}
