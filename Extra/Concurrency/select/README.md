# The select statement

> The select statement lets a goroutine wait on multiple communication operations.

A select blocks until one of its cases can run, then it executes that case. It
chooses one at random if multiple are ready.

There is a `default` case as well.

> The default case in a select is run if no other case is ready.

## Code Review

* [Example 1](example1/main.go)
* [Example 2](example2/main.go)



## Spec

* https://go.dev/ref/spec#Select_statements

> If one or more of the communications can proceed, a single one that can
> proceed is chosen via a uniform pseudo-random selection.