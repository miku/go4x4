package main

import (
	"fmt"
	"strings"

	"github.com/fatih/structs"
)

// f handles defaults hard coded within a function, caller needs to pass in the
// zero value.
func f(a, b string) string {
	if a == "" {
		a = "hello"
	}
	if b == "" {
		b = "world"
	}
	return fmt.Sprintf("%s, %s", a, b)
}

// g uses variadic final parameters, which when empty will be replaced by a
// default value.
func g(a string, b ...string) string {
	if len(b) == 0 {
		b = []string{"world"}
	}
	return fmt.Sprintf("%s, %s", a, strings.Join(b, ", "))
}

type Opts struct {
	A string `default:"hello"`
	B string `default:"world"`
}

func withDefaults(v interface{}) {
	s := structs.New(v)
	for _, f := range s.Fields() {
		if f.IsZero() && f.Tag("default") != "" {
			f.Set(f.Tag("default"))
		}
	}
}

// h accepts opts (or nil) and will fill in the defaults from the struct tag.
func h(opts *Opts) string {
	if opts == nil {
		opts = &Opts{}
	}
	withDefaults(opts)
	return fmt.Sprintf("%s, %s", opts.A, opts.B)
}

// i does not constrain the number and type of the arguments at all. It
// requires more handling and is less useful to the reader. This pattern is
// used, but rare (e.g. fmt.Printf uses it).
func i(args ...interface{}) string {
	var a, b string
	switch {
	case len(args) == 0:
		a, b = "hello", "world"
	case len(args) == 1:
		a = fmt.Sprintf("%v", args[0])
		b = "world"
	default:
		a = fmt.Sprintf("%v", args[0])
		b = fmt.Sprintf("%v", args[1])
	}
	return fmt.Sprintf("%s, %s", a, b)
}

func main() {
	var result string

	result = f("", "")
	fmt.Println(result)

	result = g("hello")
	fmt.Println(result)

	result = h(nil)
	fmt.Println(result)
	result = h(&Opts{B: "berlin"})
	fmt.Println(result)

	result = i()
	fmt.Println(result)
	result = i("hi")
	fmt.Println(result)
	result = i("hello", "berlin")
	fmt.Println(result)
	result = i("hello", 123)
	fmt.Println(result)
}
