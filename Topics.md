# Go 4x4

> Getting Started with Go

# Getting Started with Go

* 4x 4h blocks

## Block 1/4

* Intro and Motivation
* Hello World
* Role of the Go tool
    * Most used subcommands
* Go execution model
* Variables and basic types
* Slices
* Maps
* Struct Types
    * Struct Embedding
* Pointers
    * Value and Pointer receivers
* Functions
* Control structures
    * for
    * if
    * switch

## Block 2/4

* Interfaces
    * Structure Typing
    * Prefer small interfaces
    * I/O interfaces
* Go-style object orientation
    * Data and behaviour
    * Composition
* Error handling
    * Basics
    * Package errors
    * Aggregation example
    * Tracebacks
* Concurrency primitives
    * Goroutines
    * Channels
    * The select statement
* Generics
    * Function declaration
    * Function call
    * Generic struct type
    * Type constraints
    * Builtin contraints
    * Example: A generic set implementation
* Standard library highlights
    * net and net/http
    * db and db/sql

## Block 3/4

* Go project layout
    * cmd subfolder
    * grouping files
    * special names (tests, data files, ...)
* Go modules
    * Lifecycle
    * go.mod and go.sum
    * import compatibility rule
* Working with files
    * Package os
    * Package bufio
* Working with data formats: CSV
    * Reading lines
* Working with data formats: JSON
    * Struct generators
* Working with data formats: XML
    * Struct generators
* Database access with standard lib and sqlx
    * db/sql
    * sqlx
    * alternatives
* Testing
    * builtin test framework
    * testify example
* Benchmarking
    * writing benchmarks
* Fuzz Testing
    * fuzz testing code
    * limitatins
* Writing performant code
    * track memory allocations
    * use io streams

## Block 4/4

* Concurrency Utilities and Patterns
    * sync.Pool
    * sync.Once
    * sync/x/errgroup
    * sync/x/singleflight
* Writing command line tools
    * flags
    * command line arguments
    * TUI
* Shipping Go (cross-compilation, containers, ...)
    * compiler flags
    * cross-compilation
    * building containers
* Go Gotchas
* Wrap-Up
    * Discussion and Q+A

