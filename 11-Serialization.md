# Serialization

Go uses so called [struct tags](https://go.dev/ref/spec#Struct_types).

> A field declaration may be followed by an optional string literal tag. [...]
> The tags are made visible through a reflection interface and take part in type
> identity for structs but are otherwise ignored. 

## JSON

[embedmd]:# (x/encjson/main.go)

There are a few drop-in replacements for `encoding/json` which perform a bit better:

* [https://github.com/segmentio/encoding](https://github.com/segmentio/encoding)
* [https://github.com/bytedance/sonic](https://github.com/bytedance/sonic)

## XML

Note that XML mapping to a struct has flaws: [https://pkg.go.dev/encoding/xml#pkg-note-BUG](https://pkg.go.dev/encoding/xml#pkg-note-BUG)

> Mapping between XML elements and data structures is inherently flawed: an XML
> element is an order-dependent collection of anonymous values, while a data
> structure is an order-independent collection of named values.

[embedmd]:# (x/encxml/main.go)

## Generate struct from JSON or XML

Two examples from a set a couple of existing tools:

* [JSONGen](https://github.com/bemasher/JSONGen)
* [zek](https://github.com/miku/zek), try [online](https://www.onlinetool.io/xmltogo/)

