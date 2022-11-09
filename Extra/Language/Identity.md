# Identity

A few ways to check for identity:

* `==` (comparable, [ref/spec](https://go.dev/ref/spec#Comparison_operators))
* `bytes.Equal` (compares two byte slices)
* `reflect.DeepEqual`  (recursive comparison)

* additional library, e.g. [https://github.com/google/go-cmp](https://github.com/google/go-cmp)

