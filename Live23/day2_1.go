package main

import "log"

func main() {
	stats := map[string]int{
		"ok":     120,
		"failed": 2,
	} // literal init

	// create a map
	var s1 map[string]int // nil
	var s2 = map[string]int{}
	var s3 = make(map[string]int)
	s2["x"] = 1
	s2["z"] = 3
	s2["a"] = 5
	s2["b"] = 7
	delete(s2, "b")

	log.Println(s1, s2, s3)
	log.Println(stats)

	// access
	log.Println(s2["x"])
	v := s2["x"]
	log.Println(v)

	v, ok := s2["y"]
	log.Println(v, ok)

	// iteration
	for k, v := range s2 {
		log.Printf("%s => %v", k, v)
	}

	// slice iteration
	var slice1 = []string{"a", "b", "c"}
	for _, v := range slice1 {
		log.Println(v)
	}

}
