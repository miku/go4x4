# Context

> Package context defines the Context type, which carries deadlines,
> cancellation signals, and other request-scoped values across API boundaries
> and between processes.

* example use case: longer running api calls, like [sql#Conn.QueryContext](https://pkg.go.dev/database/sql#Conn.QueryContext)

In case you want to receive a context, you most likely want an `select`
statement to orchestrate results generate by the function and any cancellation
signal.

```go
ctx := context.Background()
resultsCh := make(chan *WorkResult)

// ...

for {
	select {
	case <- ctx.Done():
		// The context is over, stop processing results
		return
	case result := <- resultsCh:
		// Process the results received
	}
}
```

## Creating a context

Background is a starting context, other context can be derived. TODO is a
placeholder.

* func Background() Context
* func TODO() Context
* func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
* func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
* func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

For carrying information across boundaries.

* func WithValue(parent Context, key, val any) Context

## Notes

* context needs to be handled by the method that accepts it
* context-aware methods have been added later (e.g. in database/sql)
* some non-context-aware types are deprecated (e.g. https://pkg.go.dev/database/sql/driver#Execer)

Example usage in http server shutdown in standard library:

> [https://cs.opensource.google/go/go/+/refs/tags/go1.19.3:src/net/http/server.go;l=2808-2818](https://cs.opensource.google/go/go/+/refs/tags/go1.19.3:src/net/http/server.go;l=2808-2818)

* if added to an api, if should be the first argument
* a value that is passed around (not stored)
* best for library functions that may block or be long running
* only use Value with request-scoped variables (and sparingly, in general - it can obscure program flow)

## Examples

* [ctxadd](../ExtraSolutions/ctxadd/main.go)
* [ctxclient](../ExtraSolutions/ctxclient/main.go)
* [ctxserve](../ExtraSolutions/ctxserve/main.go)

Usage in hedged request example:

* [hedged](../ExtraSolutions/hedged/main.go)