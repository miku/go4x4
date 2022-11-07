# Project layout

## Notions

A couple notions play a role:

* a package is a directory of code (it will be named after the directory)
* one package can have many files
* one projects can have many subpackages
* there a special `internal` subpackage
* prefer short names, without stutter
* entry points live in package main, if you have many of them, shard them under a directory, typically called `cmd`  

A couple of **non-rules**:

* one struct, one file
* one package, one file
* try to avoid too generic, meaningless names, like `utils.go`, use
  `stringutils.go` or something more speaking
 
## Dependency Management

* use go modules, start with `go mod init <name>`

## Basic projects

Keep it simple, orient on business domain, rather than implementation. Group
functions and types sensibly, there is no need for a single entity per file
approach. On the other, putting everything into a single file is probably an
anti-pattern. Keep it testable.

* Have a README
* Have a `cmd/` directory for executables
* Use `internal/` directories, if you need to share code (exported identifiers), but only internally

## Minimalistic tool

Example for a minimalistic tool.

```
$ tree -sh
.
├── [4.0K]  cmd
│   └── [4.0K]  solrdump
│       └── [3.0K]  main.go
├── [  89]  go.mod
├── [ 849]  go.sum
├── [4.0K]  images
│   └── [8.0M]  8e4zf1ryf2gusi3usv329btt8.gif
├── [ 34K]  LICENSE
├── [ 725]  Makefile
├── [4.0K]  packaging
│   ├── [4.0K]  deb
│   │   └── [4.0K]  solrdump
│   │       └── [4.0K]  DEBIAN
│   │           └── [ 239]  control
│   ├── [ 227]  Makefile
│   ├── [ 718]  PKGBUILD
│   └── [4.0K]  rpm
│       ├── [ 814]  buildrpm.sh
│       └── [ 729]  solrdump.spec
└── [3.5K]  README.md

8 directories, 12 files
```

## Small tool

Example for a small tool with a couple of tests (excerpt).

```
.
├── [4.0K]  cmd
│   └── [4.0K]  zek
│       └── [3.6K]  main.go
├── [4.0K]  docs
│   ├── [ 54K]  114391.png
│   └── [5.6K]  zek.md
├── [  97]  go.mod
├── [ 577]  go.sum
├── [ 281]  io.go
├── [ 34K]  LICENSE
├── [1.2K]  Makefile
├── [5.0K]  node.go
├── [ 16K]  node_test.go
├── [ 16K]  README.md
├── [ 837]  stack.go
├── [ 607]  stack_test.go
├── [8.5K]  structwriter.go
├── [3.7K]  structwriter_test.go
├── [4.0K]  testdata
│   ├── [ 460]  w.10.go
│   ├── [ 119]  w.10.xml
│   ├── [ 477]  w.11.go
│   ├── [ 131]  w.11.xml
│   ├── [ 248]  w.12.go
│   ├── [  57]  w.12.xml
│   ├── [ 245]  w.13.go
│   ├── [  57]  w.13.xml
│   ├── [ 119]  w.1.go
│   ├── [   7]  w.1.xml
│   ├── [ 182]  w.2.go
│   ├── [  14]  w.2.xml
│   ├── [ 384]  w.3.go
│   ├── [ 189]  w.3.xml
│   ├── [2.1K]  w.4.go
│   ├── [728K]  w.4.xml
│   ├── [ 153]  w.5.go
│   ├── [  14]  w.5.xml
│   ├── [ 152]  w.6.go
│   ├── [  14]  w.6.xml
│   ├── [ 215]  w.7.go
│   ├── [  20]  w.7.xml
│   ├── [ 192]  w.8.go
│   ├── [  21]  w.8.xml
│   ├── [ 250]  w.9.go
│   └── [ 137]  w.9.xml
└── [  62]  version.go

4 directories, 42 files
```

## Example: Reverse proxy directory layout

* 63 directories, 168 files
* web folder for frontend code
* code and test code side by side
* documentation in docs
* commands id `cmd/`

Think of the user, if your write library code. How would a user import your code?

Some projects include Makefiles, although the Go tool should take care of most
tasks, that Makefile solve. However, Makefiles can be useful, too:

* include additional tasks, not related to Go
* single entry point for your project, which allows to simplify building: clone, make, done.

```
$ tree -d -I vendor
.
├── assets
│   ├── frpc
│   │   ├── static
│   │   └── statik
│   └── frps
│   ├── event
│   ├── health
│   └── proxy
├── cmd
│   ├── frpc
│   │   └── sub
│   └── frps
├── conf
│   └── systemd
├── doc
│   └── pic
├── models
│   ├── config
│   ├── consts
│   ├── errors
│   ├── msg
│   ├── nathole
│   ├── plugin
│   └── proto
│       └── udp
├── server
│   ├── controller
│   ├── group
│   ├── ports
│   ├── proxy
│   └── stats
├── tests
│   ├── ci
│   │   └── health
│   ├── config
│   ├── consts
│   ├── mock
│   └── util
├── utils
│   ├── log
│   ├── metric
│   ├── net
│   ├── util
│   ├── version
│   └── vhost
└── web
    ├── frpc
    │   └── src
    │       ├── assets
    │       ├── components
    │       ├── router
    │       └── utils
    │           └── less
    └── frps
        └── src
            ├── assets
            ├── components
            ├── router
            └── utils
                └── less

63 directories
```

## Links

* https://talks.golang.org/2014/names.slide
* https://golang.org/doc/effective_go.html#names
* https://blog.golang.org/package-names
* https://github.com/golang/go/wiki/CodeReviewComments
* [Style guideline for Go packages (rakyll/JBD)](https://rakyll.org/style-packages/)

----

## Task: Project "geomimg"

