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

