// Examples for contructing types.
package main

import "fmt"

// Constructor variants
// ====================

func NewStoreName(name string) *Store {
	return &Store{
		Name: name,
	}
}

func NewStoreKind(kind string) *Store {
	return &Store{
		Kind: kind,
	}
}

// Options
// =======

type Opts struct {
	Name string
	Kind string
}

func New(opts *Opts) *Store {
	return &Store{
		Name: opts.Name,
		Kind: opts.Kind,
	}
}

// Functional options
// ==================

// A Store we would like to configure.
type Store struct {
	Name string
	Kind string
}

// 1. Define an option type.
type Options func(*Store)

func WithName(name string) Options {
	return func(s *Store) {
		s.Name = name
	}
}

func WithKind(kind string) Options {
	return func(s *Store) {
		s.Kind = kind
	}
}

func NewStore(opts ...Options) *Store {
	s := &Store{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func main() {
	s := NewStore(
		WithName("my store"),
		WithKind("supermarket"),
	)
	fmt.Printf("%+v\n", s)
}
