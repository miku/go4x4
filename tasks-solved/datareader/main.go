package main

import (
	"fmt"
	"log"
	"math/rand"
)

type DataReader interface {
	Read() string
}

type StringReader struct {
	s string
}

func NewStringReader(s string) *StringReader {
	return &StringReader{s: s}
}

func (r *StringReader) Read() string {
	return r.s
}

type RandomReader struct{}

func (r *RandomReader) Read() string {
	return randomString(10)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	// var dr DataReader = NewStringReader("hello")
	var dr DataReader = &RandomReader{}
	fmt.Println(dr.Read())

	switch dr.(type) {
	case *StringReader:
		log.Println("used a string reader")
	case *RandomReader:
		log.Println("used a random reader")
	default:
		log.Println("unexpected type")
	}

}
