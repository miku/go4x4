# History

* First presentation in 10 Nov, 2009 ([Google Tech Talk](https://www.youtube.com/watch?v=rKnDgT73v8s))
* [Go 1](https://blog.golang.org/go1),  28 March 2012
* [Go 1.5](https://golang.org/doc/go1.5#c), no more C

Release history:
[https://golang.org/doc/devel/release](https://golang.org/doc/devel/release),
latest as of 04/2023: [go1.20.3](https://go.dev/doc/devel/release#go1.20).


## Repo History

```
$ git summary | head -30

 project  : go
 repo age : 51 years
 active   : 5136 days
 commits  : 56101
 files    : 12216
 authors  : 
  7184  Russ Cox                                                    12.8%
  4104  Robert Griesemer                                            7.3%
  2987  Rob Pike                                                    5.3%
  2451  Ian Lance Taylor                                            4.4%
  2369  Brad Fitzpatrick                                            4.2%
  1625  Austin Clements                                             2.9%
  1551  Matthew Dempsky                                             2.8%
  1496  Josh Bleecher Snyder                                        2.7%
  1435  Keith Randall                                               2.6%
  1192  Andrew Gerrand                                              2.1%
  1154  Bryan C. Mills                                              2.1%
  1026  Cherry Zhang                                                1.8%
   922  Shenghou Ma                                                 1.6%
   777  Alex Brainman                                               1.4%
   752  Mikio Hara                                                  1.3%
   690  Dmitriy Vyukov                                              1.2%
   530  Tobias Klauser                                              0.9%
   524  Adam Langley                                                0.9%
   520  Cuong Manh Le                                               0.9%
   508  Ken Thompson                                                0.9%
   488  Than McIntosh                                               0.9%
   476  Dave Cheney                                                 0.8%
   467  Nigel Tao                                                   0.8%

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

The current release is [Go 1.20](https://go.dev/blog/go1.20) from 2023-02-01.

A quick repo analysis:

* [Slides.md#go-repo-analysis](https://github.com/golang-leipzig/state-of-go-in-2022/blob/main/Slides.md#go-repo-analysis)

Repository numbers:

```
$ tokei -o json src | jq .Go.code
1681374

$ for d in $(find . -maxdepth 1 -type d | sort); do echo -n "$(basename $d) "; tokei -o json $d | jq -rc .Go.code; done | column -t | sort -k2,2 -nr
.          1681374
cmd        836753
syscall    135182
runtime    100298
net        92448
vendor     73402
crypto     69691
go         51478
internal   44097
encoding   34873
math       28432
os         19693
debug      16493
reflect    16186
image      13431
html       11112
unicode    10907
archive    10474
text       9680
time       9224
database   8863
compress   8109
testing    7635
regexp     7304
strconv    7201
sync       6605
fmt        5699
log        5320
strings    5215
bytes      5059
mime       4784
path       4314
io         3357
bufio      3151
hash       2938
sort       2511
index      2196
flag       1584
context    1581
container  1126
expvar     778
slices     737
errors     560
embed      430
maps       208
plugin     127
arena      64
builtin    52
unsafe     12

$ cd cmd
$ for d in $(find . -maxdepth 1 -type d | sort); do echo -n "$(basename $d) "; tokei -o json $d | jq -rc .Go.code; done | column -t | sort -k2,2 -nr
.          836753
compile    393499
vendor     249327
internal   71665
go         55063
link       31956
asm        6219
cgo        5083
dist       4244
trace      4103
fix        2829
doc        2497
cover      2368
covdata    1897
api        1488
vet        1074
gofmt      1065
pack       672
pprof      500
nm         467
objdump    416
addr2line  186
test2json  72
buildid    63
```

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
