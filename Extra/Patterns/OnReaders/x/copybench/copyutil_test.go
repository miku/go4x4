package copybench

import (
	"fmt"
	"io"
	"io/ioutil"
	"testing"
)

const (
	K = 1024
	M = 1024 * K
	G = 1024 * M
)

func TestSimple(t *testing.T) {
	size := 1 * M
	f := EmptyFile(size)
	k, err := io.Copy(ioutil.Discard, f)
	if err != nil {
		t.Errorf("want %v, got %v", nil, err)
	}
	if k != int64(size) {
		t.Errorf("want %d, got %d", size, k)
	}
}

func BenchmarkSimple(b *testing.B) {
	sizes := []int{1 * M, 10 * M, 100 * M, 1 * G, 10 * G, 50 * G}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(sb *testing.B) {
			f := EmptyFile(size)
			for i := 0; i < sb.N; i++ {
				_, err := io.Copy(ioutil.Discard, f)
				if err != nil {
					b.Errorf("want %v, got %v", nil, err)
				}
			}
		})
	}
}

func BenchmarkCopyBuffer(b *testing.B) {
	sizes := []int{1 * M, 10 * M, 100 * M, 1 * G, 10 * G, 50 * G}
	bufs := []int{1, 8, 1024, 4096, 8192, 16384, 32768}
	for _, size := range sizes {
		for _, bs := range bufs {
			b.Run(fmt.Sprintf("size-%d-buf-%d", size, bs), func(sb *testing.B) {
				f := EmptyFile(size)
				buffer := make([]byte, bs)
				for i := 0; i < sb.N; i++ {
					_, err := io.CopyBuffer(ioutil.Discard, f, buffer)
					if err != nil {
						b.Errorf("want %v, got %v", nil, err)
					}
				}
			})
		}
	}
}
