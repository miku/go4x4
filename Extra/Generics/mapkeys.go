package main

import "log"

// MapKeys returns a list of keys for any map type.
func MapKeys[K comparable, V any](m map[K]V) []K {
	var ks []K
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}

func main() {
	vocalMap := map[string]bool{
		"a": true,
		"b": false,
		"c": false,
		"d": false,
		"e": true,
	}
	keys := MapKeys(vocalMap)
	log.Println(keys)
}

// $ go run mapkeys.go
// 2022/11/07 21:55:07 [a b c d e]
