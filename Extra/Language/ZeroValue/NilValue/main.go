package main

import "log"

type Node struct {
	Name string
}

func (n *Node) String() string {
	if n == nil {
		return "undefined"
	}
	return n.Name
}

func main() {
	var node *Node
	log.Println(node)
}
