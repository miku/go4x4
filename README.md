# Go Workshop Notes

> 4x 4h sessions

> instructor: [about](https://github.com/miku/go4x4/blob/main/About.md)

## Outline


* [Introduction](01-Intro.md)
    * Motivation and Landscape
    * Hello World
    * Task: "helloworld"
* [History](02-History.md)
    * Timeline of events
* [More First Programs](03-MoreFirstPrograms.md)
    * Entry point
    * Importing Code
    * Visibility (public, private)
* [Language Overview](04-Language.md)
    * Basic Types
    * Variable Declarations
    * Control Structures: if, for, switch
* [More Types](05-MoreTypes.md)
    * Slices
    * Maps
* [Interfaces](06-Interfaces.md)
    * Structural Typing
    * Small interfaces
    * Variants (basic, embedding, general, [ref/spec](https://go.dev/ref/spec#Interface_types))
* [Go OOP?](07-OO.md)
    * Is Go object oriented? [FAQ](https://go.dev/doc/faq#Is_Go_an_object-oriented_language)
* [Error Handling](08-Errors.md)
    * Custom Error Types
    * Wrapping and unwrapping errors
* [Project Layout](09-Projects.md)
    * typical structure
    * naming recommendations
    * import path and resolution
    * Go modules
    * mixing public and private code
    * versioning libraries
* [IO](10-IO.md)
    * working with files
    * readers and writers
* [Serialization](11-Serialization.md)
    * struct tags
    * JSON
    * XML
* [Testing Go Code](12-Testing.md)
    * Unit Test
    * Subtests
    * Benchmarks
    * Testcontainers
* [Concurrency](13-Concurrency.md)
    * classic and CSP style
    * goroutines
    * channels
    * select
    * the `sync` package
    * error handling
    * helpers: `errgroup`
* [HTTP clients](14-HTTP.md)
    * clients and transport
    * standard clients, third party clients
* [HTTP servers](15-Servers.md)
    * handlers
    * router, e.g. gorilla/mux
    * testing
* [Database access](16-Databases.md)
    * package db, db/sql
    * database drivers
    * sqlx helper


----

## Tasks

* [tasks/README.md](tasks/README.md)

---

* [asciitable](tasks/asciitable)
* [btcprice](tasks/btcprice)
* [datareader](tasks/datareader)
* [generator](tasks/generator)
* [genreverse](tasks/genreverse)
* [helloworld](tasks/helloworld)
* [infreader](tasks/infreader)
* [linkchecker](tasks/linkchecker)
* [randomimg](tasks/randomimg)
* [rot13reader](tasks/rot13reader)
* [switch](tasks/switch)
* [timer](tasks/timer)
* [vardecl](tasks/vardecl)
* [wordfreq](tasks/wordfreq)

----

Follow up:

* Linkchecker
* Project layout, modules
    * add library
    * add service
* Services
* Misc
