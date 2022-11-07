# Common Mistakes


## Goroutines and closures

* goroutines and closures, https://golang.org/doc/faq#closures_and_goroutines

```go
func main() {
    done := make(chan bool)

    values := []string{"a", "b", "c"}
    for _, v := range values {
        go func() {
            fmt.Println(v)
            done <- true
        }()
    }

    // wait for all goroutines to complete before exiting
    for _ = range values {
        <-done
    }
}
```

> One might mistakenly expect to see a, b, c as the output. What you'll probably
see instead is c, c, c. This is because each iteration of the loop uses the same
instance of the variable v, so each closure shares that single variable. 

Solution:

```go
    for _, v := range values {
        go func(u string) {
            fmt.Println(u)
            done <- true
        }(v)
    }
```

## Shadowing variables

```go
var client *http.Client
if debug {
    client, err := CreateDebugClient()
    ...
} else {
    client, err := DefaultClient()
}
```

## JSON marshaling

* only Exported fields are serialized

if you want to suppress export of a public field, use "-"

```go
type Foo struct {
    A string
    B string
    C string `json:"-"`
}
```

## Marshaling and Embedding

```go
type Event struct {
    ID int
    time.Time
}
```

Here two things play together:

* time.Time implements `MarshalJSON` and `Event` does not
* the embedded types methods are promoted, yet the embedded type keep its
  identify

Hence, a `MarshalJSON` method is promoted on the event, but it belongs to time
only.


## Defer gotcha

Arguments to defer are evaluated right away.

Solution:

* use a pointer
* use defer with a closure

Example: [Defer](Defer/main.go)


## Not closing resources

* especially `http.Response.Body` - which is a `ReadCloser`

----

# Code style

## Nesting

* writing nested code (let the happy path flow down the left hand side)

