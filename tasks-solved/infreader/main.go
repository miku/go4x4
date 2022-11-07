package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(p []byte) (int, error) {
	for i := 0; i < len(p); i++ {
		p[i] = 'A'
	}
	return len(p), nil
}

func main() {
	reader.Validate(MyReader{})
}
