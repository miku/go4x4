package main

import (
	"log"
)

// We had some generic types before, e.g. map.
var Counter = make(map[string]int)

// List has a parameterized type, unconstrained.
type List[T any] struct {
	vs []T
}

// Append works on a List[T].
func (l *List[T]) Append(v T) {
	l.vs = append(l.vs, v)
}

func main() {
	vs := List[string]{} // List{} would err!
	log.Println(vs)
	vs.Append("a")
	vs.Append("b")
	log.Println(vs)
}

// 2022/11/07 21:48:49 {[]}
// 2022/11/07 21:48:49 {[a b]}
