# More First Programs

Let's write a few variations of our first program.

## Entry point

> The first statement in a Go source file must be package name. Executable
> commands must always use package main.  --
> [https://go.dev/doc/code](https://go.dev/doc/code)

A package `main` and a file using a different package name cannot live in the
same directory.

* example: [x/twopkgs](x/twopkgs/)

```shell
$ make
found packages main (main.go) and other (other.go) in /home/tir/code/miku/go4x4/x/twopkgs
```

This is why you see a projects with entry points separating commands into a
subdirectory, like `cmd` (which itself can be traced to
[plan9](https://github.com/0intro/plan9/tree/main/sys/src)).

Not required for one-off tools, e.g.:

* [embedmd](https://github.com/campoy/embedmd)
* [crfs](https://github.com/google/crfs)
* ...

For prototypical implementations, this minimalistic approach can be useful.

## Encoding

Go code is expected to be encoded in UTF-8.

> Source code in Go is defined to be UTF-8 text; no other representation is
> allowed. [...] To summarize, strings can contain arbitrary bytes, but when
> constructed from string literals, those bytes are (almost always) UTF-8. --
> [https://go.dev/blog/strings](https://go.dev/blog/strings)


