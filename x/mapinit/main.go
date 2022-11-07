package main

import (
	"fmt"
)

func main() {
	stats := map[string]int{
		"ok":     120,
		"failed": 2,
	}
	fmt.Println(stats)       // --> map[failed:2 ok:120]
	stats["ok"]++            // modify entry
	stats["new"] = 1         // add entry
	fmt.Println(stats)       // --> map[failed:2 new:1 ok:121]
	delete(stats, "new")     // delete an entry
	fmt.Println(stats)       // --> map[failed:2 ok:121]
	fmt.Println(stats["ok"]) // --> 121
}
