package sr

import "testing"

func BenchmarkF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f(10)
	}
}

func BenchmarkG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		g(1234.0)
	}
}

func BenchmarkH(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h(1234.0)
	}
}
