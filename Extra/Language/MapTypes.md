# Map Types

Go has a single map type, which is not thread safe. This type either uses
literals or `make` for initialization.  

Use the builtin delete to delete entries.

Possible map keys:

> The comparison operators == and != must be fully defined for operands of the
> key type; thus the key type must not be a function, map, or slice.

Maps respond to `len` and return the number of elements.

## Using a map as set

A common use of maps is to represent a set of elements. The key type is the type
of the set elements, the value type is `bool` or `struct{}`.

Example:

* [MapTypes/Set](MapTypes/Set/main.go)


## Thread safe maps

Either constuct your own with a mutex. Or use the `sync.Map` type - which is useful in specific situations only.

* [https://github.com/embano1/go-meetup-lej-04-2020/tree/master/sync/sync-map](https://github.com/embano1/go-meetup-lej-04-2020/tree/master/sync/sync-map)

![](MapTypesBenchmark.png)

