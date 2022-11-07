# Error Handling

Go has no exceptions, we say errors are values. Errors are not exceptional,
rather part of the control flow.

```go
type error interface {
    Error() string
}
```

Another example of a small interface.

Motto: Don't just check errors, handle them gracefully.

> Note: A common complaint is error handling in go and that you will write:

```go
if err := ...; err != nil {
    // handle error
}
```

over and over again. That's true. If you see an unhandled error or an `_` where
the error should be, then consider it a potential issue.

Side note: An empirical study suggested that exceptions handling is not
necessarily fine grained, or that exceptions are not handled at all ([Greg Wilson, Software Engineering's Greatest Hits](https://www.youtube.com/watch?v=HrVtA-ue-x0)).

## Ad-Hoc messages

Use [fmt.Errorf](https://pkg.go.dev/fmt#Errorf) to returns ad-hoc error messages.

Example:

* [x/ErrorHandling/AdHoc](x/ErrorHandling/AdHoc/main.go)

## Predefine errors

Use [errors.New](https://pkg.go.dev/errors#example-New) to declare errors in a package.

Advantage:

* You can check for specific errors at different call sites

The implementation of `errors.New` is simple:

```go
// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func New(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
```

Example for declared errors:

* [https://github.com/boltdb/bolt/blob/fd01fc79c553a8e99d512a07e8e0c63d4a3ccfc5/errors.go#L5-L29](https://github.com/boltdb/bolt/blob/fd01fc79c553a8e99d512a07e8e0c63d4a3ccfc5/errors.go#L5-L29)
* [https://github.com/containers/podman/blob/056f492f59c333d521ebbbe186abde0278e815db/libpod/define/errors.go#L8-L204](https://github.com/containers/podman/blob/056f492f59c333d521ebbbe186abde0278e815db/libpod/define/errors.go#L8-L204)
* [https://github.com/etcd-io/etcd/blob/e2d67f2e3bfa6f72178e26557bb22cc1482c418c/server/etcdserver/errors.go#L22-L47](https://github.com/etcd-io/etcd/blob/e2d67f2e3bfa6f72178e26557bb22cc1482c418c/server/etcdserver/errors.go#L22-L47)

Return or compare errors:

* [https://github.com/etcd-io/etcd/blob/1e46145b29764f9e65f4164bb777fbf275a63d8a/server/etcdserver/apply.go#L1062](https://github.com/etcd-io/etcd/blob/1e46145b29764f9e65f4164bb777fbf275a63d8a/server/etcdserver/apply.go#L1062)
* [https://github.com/chrislusf/seaweedfs/blob/556cc3a4ca6702346688a82e9b2fb72266c2c307/weed/util/chunk_cache/on_disk_cache_layer.go#L73-L75](https://github.com/chrislusf/seaweedfs/blob/556cc3a4ca6702346688a82e9b2fb72266c2c307/weed/util/chunk_cache/on_disk_cache_layer.go#L73-L75)

## A custom error type

The [fs.PathError](https://github.com/golang/go/blob/8f4c020660d4c8a7bab9a7363551d07176e638eb/src/io/fs/fs.go#L242-L257) is an example.

```go
// PathError records an error and the operation and file path that caused it.
type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }

// more methods ...
```


Other examples:

* [https://pkg.go.dev/encoding/csv#ParseError](https://pkg.go.dev/encoding/csv#ParseError)

## Wrapping errors, adding context

When something fails, we may want to just pass the error or **add more context** to it.

You could define a custom type, like:

```go
type ErrorStack struct {
    errs []error
}

func (s *ErrorStack) Error() string { ... }
func (s *ErrorStack) Push(err error) { ... }
```

We could also add methods to check, whether this stack contains specific errors. Two other options, which are a bit more preferred:

* use `%w` (since 1.13)
* use [https://pkg.go.dev/github.com/pkg/errors](https://pkg.go.dev/github.com/pkg/errors)

Examples:

* [ErrVerb](x/ErrorHandling/ErrVerb/main.go)
* [WrapError](x/ErrorHandling/WrapError/main.go)


## Additional Helpers

* errors.Is (compare values)
* errors.As (compare type)


```go
if errors.Is(err, ErrPermission) {
    // err, or some error that it wraps, is a permission problem
}
```

Before Go 1.13, we would have to write:

```go
if e, ok := err.(*QueryError); ok && e.Err == ErrPermission {
    // query failed because of a permission problem
}
```

Example for errors.As:

```go
var e *QueryError
// Note: *QueryError is the type of the error.
if errors.As(err, &e) {
    // err is a *QueryError, and e is set to the error's value
}
```

