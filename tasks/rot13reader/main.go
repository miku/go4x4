package main

import (
	"io"
	"os"
	"strings"
)

// 1. Define a "rot13reader" - it should be usable with any underlying reader.

func (r *rot13Reader) Read(p []byte) (n int, err error) {
	if n, err = r.r.Read(p); err != nil {
		return n, err
	}
	for i := range p {
		if 'a' <= p[i] && p[i] <= 'm' || 'A' <= p[i] && p[i] <= 'M' {
			p[i] = p[i] + 13
		} else if 'n' <= p[i] && p[i] <= 'z' || 'N' <= p[i] && p[i] <= 'Z' {
			p[i] = p[i] - 13
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
