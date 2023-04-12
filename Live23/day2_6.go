package main

import (
	"log"
	"os"
)

func ReadFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// defer func() {
	// 	log.Println("exiting")
	// 	f.Close()
	// }()
	var b = make([]byte, 10)
	_, _ = f.Read(b)
	log.Println(string(b))
	return nil
}

func main() {
	ReadFile("day2_6.go")
}
