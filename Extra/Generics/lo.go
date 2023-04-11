package main

import (
	"log"

	"github.com/samber/lo"
)

func main() {
	// names := lo.Uniq[string]([]string{"Samuel", "John", "Samuel"})
	names := lo.Uniq([]string{"Samuel", "John", "Samuel"})
	log.Println(names)
}
