# Services

## Handler

There are a couple of concepts related to services.

* [http.Handler](https://golang.org/pkg/net/http/#Handler)

The core interface for handline HTTP request.

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

## Handlerfunc

You can turn a function (with the right signature) into a type, that satisfies
http.Handler.

* [http.HandlerFunc](https://golang.org/pkg/net/http/#HandlerFunc)

> The HandlerFunc type is an adapter to allow the use of ordinary functions as
> HTTP handlers. If f is a function with the appropriate signature,
> HandlerFunc(f) is a Handler that calls f.

With first class functions, it is possible to start an HTTP handler with a
single (long) line.

## HandleFunc

If we type `http.HandleFunc`, we are working on the `DefaultServeMux`, a
[http.ServeMux](https://golang.org/pkg/net/http/#ServeMux) added for convenience
in the http package.

> ServeMux is an HTTP request multiplexer. It matches the URL of each incoming
> request against a list of registered patterns and calls the handler for the
> pattern that most closely matches the URL.

The `ServeMux` is itself an HTTP handler, as it implements
[ServeHTTP](https://golang.org/pkg/net/http/#ServeMux.ServeHTTP).


## Code Review

* [Example 1](00-server/main.go)
* [Example 2](01-server-handlerfunc/main.go)
* [Example 3](02-server-handler/main.go)
* [Example 4](03-server-handle)
* [Example 4](04-server-handlefunc/main.go)
* [Example 5](05-two-handlers/main.go)
* [Example 6](06-handler-methods/main.go)
* [Example 7](07-methods/main.go)
* [Example 8](08-post/main.go)
* [Example 9](09-pattern/main.go)
* [Example 10](10-mux/main.go)
* [Example 11](11-gorilla-mux/main.go)

## Exercises

* [Debug Handler](template1/main.go) ([Answer](exercise1/main.go))

## Project

A todo web application.


* [todosvc](https://github.com/miku/todosvc)
