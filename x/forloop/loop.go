package main

import (
	"fmt"
	"log"
)

func main() {

	// Loop control.
	var k int32
	for {
		k++
		if k == 2_147_483_647 {
			break
		}
	}
	log.Printf("%d %d", k, k+1)

	// Indexed iteration.
	for i := 0; i < 3; i++ {
		log.Printf("i=%d", i)
	}

	// A slice.
	var abc = []string{"a", "b", "c"}

	for i, c := range abc {
		log.Printf("%d %s", i, c)
	}

	// A map.
	m := map[string]string{
		"name":    "go",
		"version": "1.16",
	}
	// Garbled iteration order.
	for k, v := range m {
		fmt.Printf("%s => %s\n", k, v)
	}
}

// 2021/06/11 00:33:39 2147483647 -2147483648
// 2021/06/11 00:33:39 i=0
// 2021/06/11 00:33:39 i=1
// 2021/06/11 00:33:39 i=2
// 2021/06/11 00:33:39 0 a
// 2021/06/11 00:33:39 1 b
// 2021/06/11 00:33:39 2 c
// name => go
// version => 1.16
