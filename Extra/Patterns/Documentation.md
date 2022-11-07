# Documentation

Documentation rules are kept simple.

```go
// Package greetings groups utility functions for salutations.
package greetings

// Hello prints a greeting.
func Hello() {
    fmt.Println("Hello")
}
```


Documentation on public code will show up on [godoc.org](godoc.org), for
example:
[https://godoc.org/go4.org/readerutil#Size](https://godoc.org/go4.org/readerutil#Size).

If installed locally, the `go doc` command will print out the documetation as well.

```shell
$ go doc go4.org/readerutil
package readerutil // import "go4.org/readerutil"

Package readerutil provides and operates on io.Readers.

Package readerutil contains io.Reader types.

func NewBufferingReaderAt(r io.Reader) io.ReaderAt
func NewFakeSeeker(r io.Reader, size int64) io.ReadSeeker
func NewStatsReadSeeker(v *expvar.Int, rs io.ReadSeeker) io.ReadSeeker
func NewStatsReader(v *expvar.Int, r io.Reader) io.Reader
func Size(r io.Reader) (size int64, ok bool)
type CountingReader struct{ ... }
type ReadSeekCloser interface{ ... }
type ReaderAtCloser interface{ ... }
type SizeReaderAt interface{ ... }
    func NewMultiReaderAt(parts ...SizeReaderAt) SizeReaderAt
```

Given a package and an identifier, details for that identifier are shown.

```shell
$ go doc go4.org/readerutil.Size
package readerutil // import "go4.org/readerutil"

func Size(r io.Reader) (size int64, ok bool)
    Size tries to determine the length of r. If r is an io.Seeker, Size may seek
    to guess the length.
```

Another example from the standard library:

```shell
$ go doc http.Response.Location
package http // import "net/http"

func (r *Response) Location() (*url.URL, error)
    Location returns the URL of the response's "Location" header, if present.
    Relative redirects are resolved relative to the Response's Request.
    ErrNoLocation is returned if no Location header is present.
```

## A doc.go file

If package documentation gets longer or does not fit into any particular file,
it can by convention be separated into a `doc.go` file, which contains only the
package name and documentation.

* [src/fmt/doc.go](https://golang.org/src/fmt/doc.go)

