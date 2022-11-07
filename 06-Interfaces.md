# Interfaces

Interfaces provide a way to declare types that define only behavior. This
behavior can be implemented by concrete types, such as struct or named types,
via methods.

When a concrete type implements the set of methods for an
interface, values of the concrete type can be assigned to variables of the
interface type. Then method calls against the interface value actually call into
the equivalent method of the concrete value. Since any concrete type can
implement any interface, method calls against an interface value are polymorphic
in nature.

## Notes

* The method set for a value, only includes methods implemented with a value receiver.
* The method set for a pointer, includes methods implemented with both pointer and value receivers.
* Methods declared with a pointer receiver, only implement the interface with pointer values.
* Methods declared with a value receiver, implement the interface with both a value and pointer receiver.
* Interfaces are reference types, don't share with a pointer.
* This is how we create polymorphic behavior in go.

## Examples

* [x/interfaces/main.go](x/interfaces/main.go)

## Gotcha

Example nil interface:

* [x/interfacenil/main.go](x/interfacenil/main.go)

## Type Assertion and Type Switch

* type assertions can retrieve the concrete type
* type switch allows to switch based on type

## Interface Composition

* [https://pkg.go.dev/io#ReadWriter](https://pkg.go.dev/io#ReadWriter)

```
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}
```