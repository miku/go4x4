# Language Features

## Variables

* There are four notations to
  [declare](https://golang.org/ref/spec#Variable_declarations) and initialize
variables.

You'll mostly the following two (type inference):

[embedmd]:# (x/vars/main.go /^.*a := 3.14/ /var b = 1.0.*/)
```go
	a := 3.14        // assign and infer type (only within functions)
	var b = 1.0      // assign and infer type (anywhere)
```

You can name the type explicitly. The type follows the identifier (as
opposed to C, Java, ...):

[embedmd]:# (x/vars/main.go /^.*c int/ /var d int8 =.*/)
```go
	var c int        // explicit type, without assignment
	var d int8 = 127 // explicit type, with assignment
```

* The default integer type will by [int](https://play.golang.org/p/z1emUwkp1rL).
* Every type has a [zero value](https://golang.org/ref/spec#The_zero_value)

> **When storage is allocated for a variable**, either through a declaration or
> a call of new, or when a new value is created, either through a composite
> literal or a call of make, **and no explicit initialization is provided, the
> variable or value is given a default value**. Each element of such a variable
> or value is set to the **zero value for its type**: false for booleans, 0 for
> numeric types, "" for strings, and nil for pointers, functions, interfaces,
> slices, channels, and maps. This initialization is done recursively, so for
> instance each element of an array of structs will have its fields zeroed if
> no value is specified.

## Types

### Primitive Types

* boolean, called `bool`, predeclared constants `true` and `false`
* [https://golang.org/ref/spec#Numeric_types](https://golang.org/ref/spec#Numeric_types)


[embedmd]:# (https://golang.org/ref/spec go /^uint8.*/ /^rune.*/)
```go
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
```


There are implementation specific predeclared types for `uint` (32 or 64 bits),
`int` (same as `uint`), `uintptr` large enough to store pointer value.

* [proposal: spec: remove complex numbers #19921](https://github.com/golang/go/issues/19921) (Apr 11, 2017)

> Go supports complex numbers, but ~nobody uses them.

* String type, `s := "Hello 世界"` - all Go source is UTF-8 encoded, double quotes or backticks for multiline strings; strings are *immutable*.

Any UTF-8 character may be used as a variable name
([play](https://play.golang.org/p/eU4VDR0-jpE)):

[embedmd]:# (x/hellojp/main.go)
```go
package main

import (
	"fmt"
)

func main() {
	世界 := "hello world"
	fmt.Println(世界)
}
```

### Arrays and Slices

There are [arrays](https://golang.org/ref/spec#Array_types) and
[slices](https://golang.org/ref/spec#Slice_types) in Go. Arrays carry their
length as part of the type, slices are dynamically sized arrays. You will
mostly encounter slices.

Example (on [play](https://play.golang.org/p/bsgboAZ82jH)):

[embedmd]:# (x/slicehello/main.go)
```go
package main

import (
	"fmt"
)

func main() {
	data := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v has length %d", data, len(data))
}
```

#### Creating a slice

[embedmd]:# (x/sliceinit/main.go)
```go
package main

func main() {
	var a []string
	b := []string{}
	c := []string{"a", "b", "c"}
}
```

#### The append builtin function

* [Go Tour: Appending to slice](https://tour.golang.org/moretypes/15), it is a
  [built-in](https://golang.org/pkg/builtin/#append) function

[embedmd]:# (x/sliceappend/main.go)
```go
package main

func main() {
	var a []string
	a = append(a, "x")
	a = append(a, "y")

	b := []string{"X", "Y"}
	a = append(a, b...)
}
```

#### Custom slice indices

Very rare, but possible: custom indices for slice elements on literal initializion:

[embedmd]:# (x/sliceindex/main.go)
```go
package main

import "log"

func main() {

	// zero indexed, by default
	zMonths := []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}
	log.Println(len(zMonths), zMonths)
	log.Println(zMonths[0])

	// custom indices possible
	months := []string{
		3:  "mar",
		1:  "jan",
		2:  "feb",
		4:  "apr",
		5:  "may",
		6:  "jun",
		7:  "jul",
		8:  "aug",
		9:  "sep",
		10: "oct",
		11: "nov",
		12: "dec",
	}
	log.Println(len(months), months)
	log.Println(months[1])
	log.Println(months[0]) // does not panic, just returns the empty string

	// log.Println(months[13]) // panic: runtime error: index out of range [13] with length 13

}
```

### Maps

* Map types are the way to implement hash maps or associative arrays.

Example ([play](https://play.golang.org/p/kMy733hPcFT)):

[embedmd]:# (x/mapinit/main.go)
```go
package main

import (
	"fmt"
)

func main() {
	stats := map[string]int{
		"ok":     120,
		"failed": 2,
	}
	fmt.Println(stats)       // --> map[failed:2 ok:120]
	stats["ok"]++            // modify entry
	stats["new"] = 1         // add entry
	fmt.Println(stats)       // --> map[failed:2 new:1 ok:121]
	delete(stats, "new")     // delete an entry
	fmt.Println(stats)       // --> map[failed:2 ok:121]
	fmt.Println(stats["ok"]) // --> 121
}
```

* [runtime/map.go](https://github.com/golang/go/blob/2169deb35247a80794519589e7cd845c6ebf4e5a/src/runtime/map.go#L7-L54)

### Structs

* Struct types group zero or more fields (name, type) into a type.

Example ([play](https://play.golang.org/p/RYOrTC-CqPu)):

[embedmd]:# (x/structinit/main.go)
```go
package main

import "fmt"

type Location struct {
	Lat  float64
	Long float64
}

type Peer struct {
	Name     string
	Location Location
}

func main() {
	peer := Peer{
		Name: "martin",
		Location: Location{
			Lat:  51.33962,
			Long: 12.37129,
		}}
	fmt.Println(peer)
	// {martin {51.33962 12.37129}}
}
```

Struct types support embedding (allows for reuse and composition) and tags (a
weakly typed, but lightweight approach to annotations).

Example ([play](https://play.golang.org/p/KgKB3jqCapV)):

[embedmd]:# (x/structembed/main.go)
```go
package main

import (
	"encoding/json"
	"fmt"
)

type Location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type Peer struct {
	Name     string   `json:"name"`
	Location Location `json:"loc"`
}

func main() {
	peer := Peer{
		Name: "martin",
		Location: Location{
			Lat:  51.33962,
			Long: 12.37129,
		}}
	b, _ := json.Marshal(peer)
	fmt.Println(string(b))
	// {"name":"martin","loc":{"lat":51.33962,"long":12.37129}}
}
```

## Control Flow

### for

* the only loop construct

Takes various forms:

* endless loop
* index variable
* iteration over containers
* receiving from channels

[embedmd]:# (x/forloop/loop.go)
```go
package main

import (
	"fmt"
	"log"
)

func main() {

	// Loop control.
	var k int32
	for {
		k++
		if k == 2_147_483_647 {
			break
		}
	}
	log.Printf("%d %d", k, k+1)

	// Indexed iteration.
	for i := 0; i < 3; i++ {
		log.Printf("i=%d", i)
	}

	// A slice.
	var abc = []string{"a", "b", "c"}

	for i, c := range abc {
		log.Printf("%d %s", i, c)
	}

	// A map.
	m := map[string]string{
		"name":    "go",
		"version": "1.16",
	}
	// Garbled iteration order.
	for k, v := range m {
		fmt.Printf("%s => %s\n", k, v)
	}
}

// 2021/06/11 00:33:39 2147483647 -2147483648
// 2021/06/11 00:33:39 i=0
// 2021/06/11 00:33:39 i=1
// 2021/06/11 00:33:39 i=2
// 2021/06/11 00:33:39 0 a
// 2021/06/11 00:33:39 1 b
// 2021/06/11 00:33:39 2 c
// name => go
// version => 1.16
```

### if

* no braces required
* has a short form

[embedmd]:# (x/ifshort/main.go)
```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

### switch

* a powerful switch statement
* fallthrough, not "break"

[embedmd]:# (x/switch/main.go)
```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```

* Recap with the Go Tour: [https://tour.golang.org/flowcontrol/1](https://tour.golang.org/flowcontrol/1), ff.

## Functions

* multiple return values
* first class objects

See:

* [https://tour.golang.org/basics/4](https://tour.golang.org/basics/4)
* [https://tour.golang.org/basics/5](https://tour.golang.org/basics/5)
* [https://tour.golang.org/basics/6](https://tour.golang.org/basics/6)
* [https://tour.golang.org/basics/7](https://tour.golang.org/basics/7)

* `defer` allows to run code when the function exits
* function can have named return values


## Methods and Interfaces

* a new "interface type"
* defines a method set
* implicitly satisfied (structural typing)

See:

* [https://tour.golang.org/methods/1](https://tour.golang.org/methods/1), and on


## Errors

* no exceptions
* errors are values

> But remember: Whatever you do, always check your errors! --
> [https://blog.golang.org/errors-are-values](https://blog.golang.org/errors-are-values)

More in: [Errors.md](08-Errors.md)


## Extra: Go Tour More Types

* [https://tour.golang.org/moretypes/1](https://tour.golang.org/moretypes/1)


## Workout

* [Exercise: loops and functions](https://tour.golang.org/flowcontrol/8)

>  As a way to play with functions and loops, let's implement a square root
>  function: given a number x, we want to find the number z for which z² is most
>  nearly x.

> Computers typically compute the square root of x using a loop. Starting with
> some guess z, we can adjust z based on how close z² is to x, producing a
> better guess:

> z -= (z*z - x) / (2*z)

> Repeating this adjustment makes the guess better and better until we reach an
> answer that is as close to the actual square root as can be.


* [Exercise: maps](https://tour.golang.org/moretypes/23)

> Implement WordCount. It should return a map of the counts of each “word” in
> the string s. The wc.Test function runs a test suite against the provided
> function and prints success or failure.

You might find [strings.Fields](https://golang.org/pkg/strings/#Fields) helpful.

## Generics

* [Extra/Generics/](Extra/Generics/)

## Roundup and Extra Topics

* [Extra/Language](Extra/Language)

Slice gotchas:

* [Extra/Language/Slices.md](Extra/Language/Slices.md)
