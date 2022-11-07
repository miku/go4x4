# Comments

Go supports various 

## go:generate

* https://golang.org/doc/go1.4#gogenerate

Allows to generate code at build time with:

```
$ go generate
```

Use by some serialization libraries.

## go:embed

* https://pkg.go.dev/embed

Allows to embed files. Works on:

* string
* byte
* FS

```
//go:embed myfile.txt
var file string
```

Or one or more directories, which will be accessible as a file system.

```
// content is our static web server content.
//go:embed image template html/index.html
var content embed.FS
```

> If a pattern names a directory, all files in the subtree rooted at that
> directory are embedded (recursively), except that files with names beginning
> with ‘.’ or ‘_’ are excluded.


## build tags

Allows to include or exclude a file from the build process. The comment must
appear at the beginning of the file.

The build syntax is currently under review.

```
// +build tag_name
```

For example to compile on "linux" OR "darwin":

```
// +build darwin linux
```

To include on ("linux" OR "darwin") AND "amd64", one would write:

```
// +build linux darwin
// +build amd64
```

Values of GOOS and GOARCH can be used here.