# History

* First presentation in 10 Nov, 2009 ([Google Tech Talk](https://www.youtube.com/watch?v=rKnDgT73v8s))
* [Go 1](https://blog.golang.org/go1),  28 March 2012
* [Go 1.5](https://golang.org/doc/go1.5#c), no more C

Release history:
[https://golang.org/doc/devel/release](https://golang.org/doc/devel/release),
latest as of 05/2022: [go1.18.1](https://go.dev/doc/devel/release#go1.18).


## Repo History

```
$ git summary | head -30

 project  : go
 repo age : 50 years
 active   : 4802 days
 commits  : 52130
 files    : 11457
 authors  : 
  6936  Russ Cox                                                    13.3%
  3840  Robert Griesemer                                            7.4%
  2981  Rob Pike                                                    5.7%
  2359  Brad Fitzpatrick                                            4.5%
  2274  Ian Lance Taylor                                            4.4%
  1531  Austin Clements                                             2.9%
  1495  Josh Bleecher Snyder                                        2.9%
  1356  Matthew Dempsky                                             2.6%
  1306  Keith Randall                                               2.5%
  1192  Andrew Gerrand                                              2.3%
  1024  Cherry Zhang                                                2.0%
   923  Bryan C. Mills                                              1.8%
   922  Shenghou Ma                                                 1.8%
   772  Alex Brainman                                               1.5%
   752  Mikio Hara                                                  1.4%
   690  Dmitriy Vyukov                                              1.3%
   523  Adam Langley                                                1.0%
   508  Ken Thompson                                                1.0%
   476  Dave Cheney                                                 0.9%
   464  Nigel Tao                                                   0.9%
   418  Joel Sing                                                   0.8%
   416  Tobias Klauser                                              0.8%
   392  David Crawshaw                                              0.8%

```

First commit:

```
commit 7d7c6a97f815e9279d08cfaea7d5efb5e90695a8
Author: Brian Kernighan <bwk>
Date:   Tue Jul 18 19:05:45 1972 -0500
```

Tree at first commit:


```
$ tree
.
└── src
    └── pkg
        └── debug
            └── macho
                └── testdata
                    └── hello.b

5 directories, 1 file

$ cat src/pkg/debug/macho/testdata/hello.b
main( ) {
        extrn a, b, c;
        putchar(a); putchar(b); putchar(c); putchar('!*n');
}
a 'hell';
b 'o, w';
c 'orld';
```

## Go: A timeline of events

* [The Go programming language](https://www.youtube.com/watch?v=rKnDgT73v8s), November 2009
* [Go version 1 is released](https://blog.golang.org/go-version-1-is-released), 28 March 2012

The Go 1 release marked an important milestone, which is codified in [Go 1 and
the Future of Go Programs](https://golang.org/doc/go1compat):

> Go 1 defines two things: first, the specification of the language; and
> second, the specification of a set of core APIs, the "standard packages" of
> the Go library. The Go 1 release includes their implementation in the form
> of two compiler suites (gc and gccgo), and the core libraries themselves.

> It is intended that programs written to the Go 1 specification will continue
> to compile and run correctly, unchanged, over the lifetime of that
> specification.

In Spring 2016, Brad Fitzpatrick gave a talk: [Go 1.6: Asymptotically
approaching boring](https://www.youtube.com/watch?v=4Dr8FXs9aJM)
([slides](https://docs.google.com/presentation/d/1JsCKdK_AvDdn8EkummMNvpo7ntqteWQfynq9hFTCkhQ/view?slide=id.p)).

Go has releases approximately every six months (seems attractive to
[others](https://www.infoq.com/news/2017/09/Java6Month), too).

Rationale: Stable foundation - to build stuff on top.

* Language ([ref/spec](https://golang.org/ref/spec))
* Standard Library ([pkg](https://golang.org/pkg/))
* Runtime ([GC](https://www.youtube.com/watch?v=aiv1JOfMjm0), scheduler, and
  other pieces under active development)
* Tools (go, godoc, go vet, gofmt, goimports, ... under active development)
* Ecosystem (external packages, conferences, user groups, and much more)

Further milestones:

* Go 1.13 added modules for dependency management
* Go 1.16 added file embedding 
* Go 1.18 added generic type support, automated test via fuzzing

The current release is [Go 1.18](https://blog.golang.org/go1.18) from 2022-03-15.

## Origins and influences

A multitude of ideas and influences:

* C++ might be slow to compile and bloated
* all programming languages seem to add and add features
* we entered the multicore era ([Free Lunch Is Over](http://www.gotw.ca/publications/concurrency-ddj.htm))
* design for a networked world
* bridge the gap between dynamic and static programming languages (be safe, yet
  ease to write)
* focus on long-term maintenance
* designed with tools in mind (with `gofmt` being the prototypical tool)
* a different approach to concurrency
* a stripped version of object orientation

Original designers: Robert Griesemer, Rob Pike, Ken Thomson.
