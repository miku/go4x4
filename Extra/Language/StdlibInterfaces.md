# Stdlib Interfaces

## Package io

* io.Reader
* io.Writer

# Go Proverb

* [The bigger the interface, the weaker the abstraction](https://youtu.be/PAAkCSZUG1c?t=5m18s)


> More of theses at https://go-proverbs.github.io/

----

# Exemplified in package io

Generic I/O with `io.Reader` and `io.Writer` and a few other interfaces.

> https://golang.org/pkg/io/


----

|                    | R | W | C | S |
|--------------------|---|---|---|---|
| io.Reader          | x |   |   |   |
| io.Writer          |   | x |   |   |
| io.Closer          |   |   | x |   |
| io.Seeker          |   |   |   | x |
| io.ReadWriter      | x | x |   |   |
| io.ReadCloser      | x |   | x |   |
| io.ReadSeeker      | x |   |   | x |
| io.WriteCloser     |   | x | x |   |
| io.WriteSeeker     |   | x |   | x |
| io.ReadWriteCloser | x | x | x |   |
| io.ReadWriteSeeker | x | x |   | x |

----

## Missing things

Libraries might implement missing pieces, e.g.

* [ReadSeekCloser, ReaderAtCloser](https://github.com/go4org/go4/blob/94abd6928b1da39b1d757b60c93fb2419c409fa1/readerutil/readerutil.go#L33-L43)

From: [github.com/go4org/go4](https://github.com/go4org/go4).

----

## IO interface list

Some utility interfaces, e.g. for multithreaded IO and performance optimizations.

* `io.ReaderAt` (p, off)
* `io.ReaderFrom` (r)
* `io.WriterAt` (p, off)
* `io.WriterTo` (w)

----

## Use cases | io.ReaderAt

* `io.ReaderAt`, `io.WriterAt` -- (parallel writes) with offset

Sidenote: For filesystems, there is a [pread(2) system call](http://man7.org/linux/man-pages/man2/pread.2.html) in Linux

> read from or write to a file descriptor at a given offset ...
> The pread() and pwrite() system calls are especially useful in **multithreaded applications**.  They allow multiple threads to perform I/O on the **same file descriptor** without being affected by changes to the file offset by other threads.

* HTTP [range request example](https://github.com/snabb/httpreaderat)
* Example: list archived filenames in remote zip file without download it: [`examples/rangerequest`](https://github.com/miku/io15min/blob/master/examples/rangerequest/main.go) 

----

## RFC 7233 HTTP Range Requests 

> Likewise, devices with limited local storage might benefit from being able to request only a subset of a larger representation, such as a single page of a very large document, or the dimensions of an embedded image. --https://tools.ietf.org/html/rfc7233#section-1

![](images/rangerequest.png)

----

## Use cases | io.ReaderFrom

* [Optimizing Copy](https://medium.com/go-walkthrough/go-walkthrough-io-package-8ac5e95a9fbd) 

> To avoid using an intermediate buffer entirely, types can implement interfaces to read and write directly. When implemented, the Copy() function will **avoid the intermediate buffer** and use these implementations directly.

----

## Use cases | io.ReaderFrom


```
// io.go, https://golang.org/src/io/io.go
// ...
// Similarly, if the writer has a ReadFrom method,
// use it to do the copy.

if rt, ok := dst.(ReaderFrom); ok {
	return rt.ReadFrom(src)
}
```

Also known as: [interface upgrade](http://avtok.com/2014/11/05/interface-upgrades.html).


>  The zero-copy IO in Go is so elegant.

* https://news.ycombinator.com/item?id=8714051 (174, 2014)