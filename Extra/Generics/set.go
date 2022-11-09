// Package sets implements sets of any comparable type.
package main

import "log"

// Set is a set of values.
type Set[T comparable] map[T]struct{}

// Make returns a set of some element type.
func Make[T comparable]() Set[T] {
	return make(Set[T])
}

// Add adds v to the set s.
// If v is already in s this has no effect.
func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

// Delete removes v from the set s.
// If v is not in s this has no effect.
func (s Set[T]) Delete(v T) {
	delete(s, v)
}

// Contains reports whether v is in s.
func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func main() {
	var s = Set[string]{} // underlying type is a map, could use make here as well
	s.Add("a")
	s.Add("b")
	log.Println(s)
}
