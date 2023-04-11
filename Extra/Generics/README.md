## Generics

Or: [parametric polymorphism](https://en.wikipedia.org/wiki/Parametric_polymorphism)

Relevant proposal:

* [43651](https://github.com/golang/proposal/blob/master/design/43651-type-parameters.md) (2021, accepted)

Previous discussion:

> Go should support some form of generic programming.
Generic programming enables the representation of algorithms and data
structures in a generic form, with concrete elements of the code (such as
types) factored out. It means the ability to express algorithms with minimal
assumptions about data structures, and vice-versa.

[...]

> Generics permit type-safe polymorphic containers. Go currently has a very
> limited set of such containers: slices, and maps of most but not all types.
> Not every program can be written using a slice or map.

> Look at the functions SortInts, SortFloats, SortStrings in the sort package.
> Or SearchInts, SearchFloats, SearchStrings. Or the Len, Less, and Swap
> methods of byName in package io/ioutil. **Pure boilerplate copying**.

Who wrote that? [...] [Proposal: Go should have generics](https://github.com/golang/proposal/blob/master/design/15292-generics.md) (2011)

A few more:

> As Russ pointed out, **generics are a trade off between programmer time, compilation time, and execution time.**

### Key Ideas

* Functions and types can have **type parameters**, which are defined using **constraints**, which are **interface types**.
* Constraints describe the **methods required** and the **types permitted** for a type argument.
* Constraints describe the methods and operations permitted for a type parameter.
* **Type inference** will often permit omitting type arguments when calling functions with type parameters.

> This design is completely backward compatible.

There are many things Go Generics do not support: No metaprogramming (macros);
no operator methods (e.g. no syntax like `c[i]` for custom types); no
[covariance or contravariance](https://en.wikipedia.org/wiki/Covariance_and_contravariance_(computer_science)) of function parameters.


Notes: e.g. in Go we cannot express `List[string]` is not a subtype of
`List[any]`.

In summary:

* no specialization (no type hierarchies), no co/contra-variance
* no metaprogramming 
* no operator overloading

### Notable changes

* `any` (alias for `interface{}`, [operations
permitted](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#operations-permitted-for-any-type))
and `comparable` as [predeclared
identifiers](https://go.dev/ref/spec#Predeclared_identifiers)
* a new [contraints](https://pkg.go.dev/golang.org/x/exp/constraints) package,
not in the standard library yet (but in
[exp](https://pkg.go.dev/golang.org/x/exp#section-readme), *In short, code in
this subrepository is not subject to the Go 1 compatibility promise.*)

> What is [comparable](https://go.dev/ref/spec#Comparison_operators)?

All non-interface types which are comparable.

> In any comparison, the first operand must be assignable to the type of the
> second operand, or vice versa. 

Example:

> * Boolean values are comparable. Two boolean values are equal if they are either both true or both false.
> * Integer values are comparable and ordered, in the usual way. 

[...]

What is [assignable](https://go.dev/ref/spec#Assignability)?

> A value x of type V is assignable to a variable of type T if ...

* V and T are identical. 
* V and T have identical underlying types ...




### Syntax

```
func G[T Constraint](p T) { }
```

* contraint may be `any` ~ `interface{}` or `comparable` or an interface with methods and

### Other Bits

* interfaces allow to capture common behaviour

> In other words, interface types in Go are a form of generic programming. They
> let us capture the common aspects of different types and express them as
> methods. -- [Go generic programming today](https://go.dev/blog/why-generics)

There is a container package in the standard library, not widely used:

* [https://pkg.go.dev/container/list@go1.20.3](https://pkg.go.dev/container/list@go1.20.3)

> the Go standard library package container is mostly unused. Everything in the
> container package deals with interface{}/any values, which is Go for
> "literally anything". -- [We have Go 2](https://xeiaso.net/blog/we-have-go-2)

### Examples


#### Reverse

Finally, we can write a generic reverse for slices of **any** type.

* [https://go.dev/play/p/LDdnZMiQN6b](https://go.dev/play/p/LDdnZMiQN6b)

```go
package main

import "fmt"

// Reverse returns a new slice, with elements in reverse order.
func Reverse[T any](vs []T) []T {
    var (
        n      = len(vs)
        result = make([]T, n)
    )
    for i, v := range vs {
        result[n-1-i] = v
    }
    return result
}

func main() {
    var (
        s  = []string{"a", "b", "c"}
        i  = []int{1, 2, 3}
        f  = []float64{1, 2, 3}
        rs = Reverse(s)
        ri = Reverse(i)
        rf = Reverse(f)
    )
    fmt.Printf("%v %T\n", rs, rs)
    fmt.Printf("%v %T\n", ri, ri)
    fmt.Printf("%v %T\n", rf, rf)
    // [c b a] []string
    // [3 2 1] []int
    // [3 2 1] []float64
}
```


#### Interface Constraint

* [https://go.dev/play/p/MM38gRuTRKB](https://go.dev/play/p/MM38gRuTRKB)

```go
// Dummy example, showing contraints as interface type and type parameters.

type Number interface {
    int | int64 | float64
}

func Min[Number T](u, v T) T {
    if u < v {
        return u
    }
    return v
}
```

#### Constraints and Methods

* [https://go.dev/play/p/HsJtEJnzJdl](https://go.dev/play/p/HsJtEJnzJdl)

```go
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Number interface {
	~int | ~int64 | ~float64
	String() string
}

func Min[T Number](u, v T) T {
	if u < v {
		return u
	}
	return v
}

type MyInt int

func (i MyInt) String() string {
	return fmt.Sprintf("[%d]", i)
}

func main() {
	r := Min(MyInt(3), MyInt(4))
	fmt.Printf("%v %T", r, r)
	// 3 int
}
```

#### Set

```go
// Package sets implements sets of any comparable type.
package sets

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
```

#### Dot Product

```go
// Numeric is a constraint that matches any numeric type.
// It would likely be in a constraints package in the standard library.
type Numeric interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
        ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
        ~float32 | ~float64 |
        ~complex64 | ~complex128
}

// DotProduct returns the dot product of two slices.
// This panics if the two slices are not the same length.
func DotProduct[T Numeric](s1, s2 []T) T {
    if len(s1) != len(s2) {
        panic("DotProduct: slices of unequal length")
    }
    var r T
    for i := range s1 {
        r += s1[i] * s2[i]
    }
    return r
}
```


### Summary

* basic generic programming support
* libraries already on the way: [https://github.com/samber/lo](https://github.com/samber/lo), [https://github.com/elliotchance/pie](https://github.com/elliotchance/pie)

Examples:

```go
names := lo.Uniq([]string{"Samuel", "Marc", "Samuel"})
```

Calculating an average or finding the mode:

* [https://pkg.go.dev/github.com/elliotchance/pie/v2#Average](https://pkg.go.dev/github.com/elliotchance/pie/v2#Average)
* [https://pkg.go.dev/github.com/elliotchance/pie/v2#Mode](https://pkg.go.dev/github.com/elliotchance/pie/v2#Mode)

And more. Functional programming gets more realistic in Go:

* [https://pkg.go.dev/github.com/samber/mo#Either](https://pkg.go.dev/github.com/samber/mo#Either)
* [https://github.com/nikgalushko/fx](https://github.com/nikgalushko/fx)

Generic data structures:

* [https://github.com/zyedidia/generic](https://github.com/zyedidia/generic)

Utilities:

* [https://goldziher.github.io/go-utils/](https://goldziher.github.io/go-utils/)

Generic concurrency patterns:

* [https://github.com/lobocv/simpleflow](https://github.com/lobocv/simpleflow)

We'll see more of these this year; standard library may get some generic
support for slices and maps (like map keys, etc).

#### Other use case

* any utilities on container types
* different type need common logic (e.g. sorting a slice)

Non use cases:

* basic interface already provide a form of generic programming

If a function uses a interface (and not a type set) as a contraint, it may not
be a good use case.