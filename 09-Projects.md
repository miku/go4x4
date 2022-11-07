# Projects

Let's build a small command line project from scratch.

* create a new directory
* go mod init "calc"
* create a new file "calc.go"
    * implement a function `Add(a, b float64) float64`
* create a new directory "cmd"
* put a new file under "cmd/calc/main.go"
    * create a `main` function
    * print out the result of `calc.Add` with hard-coded values

## Layout

```
.
├── calc.go
├── cmd
│   └── calc
│       └── main.go
├── go.mod
├── LICENSE
└── README.md
```

## Documentation

Inline docs.

----

# Code and Naming Conventions

## Always format your code

Always go fmt your code. Or use goimports (superset of go fmt).

## Do not stutter

If you have a package `time` and a function `TimeSince` you should keep it brief by naming your function just `Since` so the user of your code can write:

```go
time.Since(...)
```

* [Package names](https://blog.golang.org/package-names)

## Short variable names

It is not unreasonable to follow conventions and to name an `io.Writer` just
`w`. An excerpt for `src/io/io.go`.

```go
// WriteString writes the contents of the string s to w, which accepts a slice of bytes.
// If w implements StringWriter, its WriteString method is invoked directly.
// Otherwise, w.Write is called exactly once.
func WriteString(w Writer, s string) (n int, err error) {
	if sw, ok := w.(StringWriter); ok {
		return sw.WriteString(s)
	}
	return w.Write([]byte(s))
}
```

## Package names

Try to use simple names.

> A package names with under_scores, hy-phens or mixedCaps should be avoided.

## Keep an eye on your code

The `go vet` tool is available out of the box.

> Vet examines Go source code and reports suspicious constructs, such as Printf
> calls whose arguments do not align with the format string.

It is also part of the Hackathon project, [Go Report
Card](https://goreportcard.com/), along with a few other checks.


----

# Cross compilation

Go (w/o cgo) supports cross compilation out of the box. There are two relevant environment variables:

* GOOS
* GOARCH

Example creating an ARM 32 bit Linux executable on an x86_64 Linux.

```
$ uname -sp
Linux x86_64

$ env GOOS=linux GOARCH=arm go build main.go
$ file main
main:    ELF 32-bit LSB executable, ARM, EABI5 version 1 (SYSV), statically linked, not stripped
```

To list all supported platforms (44 currently), run:

```shell
$ go tool dist list
aix/ppc64
android/386
android/amd64
android/arm
android/arm64
darwin/386
darwin/amd64
darwin/arm
darwin/arm64
dragonfly/amd64
freebsd/386
freebsd/amd64
freebsd/arm
illumos/amd64
js/wasm
linux/386
linux/amd64
linux/arm
linux/arm64
linux/mips
linux/mips64
linux/mips64le
linux/mipsle
linux/ppc64
linux/ppc64le
linux/s390x
nacl/386
nacl/amd64p32
nacl/arm
netbsd/386
netbsd/amd64
netbsd/arm
netbsd/arm64
openbsd/386
openbsd/amd64
openbsd/arm
openbsd/arm64
plan9/386
plan9/amd64
plan9/arm
solaris/amd64
windows/386
windows/amd64
windows/arm
```
