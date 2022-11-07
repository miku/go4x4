# IO

Package `io` implements interfaces and helpers.

## Files

Regular file type is [`*os.File`](https://pkg.go.dev/os#File) 

## Interfaces

* [The bigger the interface, the weaker the abstraction](https://youtu.be/PAAkCSZUG1c?t=5m18s)

> More of theses at https://go-proverbs.github.io/

Exemplified in package io: Generic I/O with `io.Reader` and `io.Writer` and a
few other interfaces.

> https://golang.org/pkg/io/


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

Libraries might implement missing pieces, e.g.

* [ReadSeekCloser, ReaderAtCloser](https://github.com/go4org/go4/blob/94abd6928b1da39b1d757b60c93fb2419c409fa1/readerutil/readerutil.go#L33-L43)

From: [github.com/go4org/go4](https://github.com/go4org/go4).

