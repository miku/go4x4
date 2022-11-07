# Empty Interface Antipattern

Go has a relatively weak type system.

> Go intentionally has a weak type system, and there are many restrictions that
> can be expressed in other languages but cannot be expressed in Go. Go in
> general encourages programming by writing code rather than programming by
> writing types. (Ian Lance Taylor --
> [https://github.com/golang/go/issues/29649#issuecomment-454820179](https://github.com/golang/go/issues/29649#issuecomment-454820179))

Also:

> Note that the cost of a feature is not the cost of implementing it, which is
> generally inconsequential (or, for some requested features, impossible), but
> the cost of forcing everybody who learns the language to learn about the
> feature. (ibid.)

* no parameterized types (generics)
* no "a function that returns the same type that it takes"

Sometimes, people start to use the empty interface `interface{}` for everything.
It is possible (and useful at time), but generally an anti-pattern.