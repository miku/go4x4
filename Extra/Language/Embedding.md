# Embedding

Go does not strive for type hierarchies, it prefers composition and interfaces
for method dispatch.

> Code reuse is not provided by type hierarchy but via composition. Language
> ecosystems with classical inheritance is often suffering from excessive level
> of indirection and premature abstractions based on inheritance which later
> makes the code complicated and unmaintainable.

Embedding types provide the final piece of sharing and reusing state and
behavior between types. Through the use of inner type promotion, an inner type's
fields and methods can be directly accessed by references of the outer type.

## Composition of interfaces

You can use embedding to combine interfaces.

Example:

```go
type ReadCloser interface {
    io.Reader
    io.Closer
}
```

Useful to start with small interfaces and and composite interfaces as you need
them.

## Notes

* Embedding types allow us to share state or behavior between types.
* The inner type never loses its identity.
* **This is not inheritance**.
* Through promotion, inner type fields and methods can be accessed through the outer type.
* The outer type can override the inner type's behavior.

## Typical Example: Embedding a lock

* [Embedding](Embedding/main.go)
