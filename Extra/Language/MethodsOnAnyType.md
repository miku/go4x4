# Methods on any type

More esoteric examples:

* overriding builtin types (e.g. `type MyString string`) and defining methods on that

Example:

* A string with a specific semantic: [embargo](https://github.com/miku/span/blob/89ccf1016797af91cefe2d2b8865743f646dedf7/licensing/embargo.go#L90)

As mentioned in [FirstClassFunctions](FirstClassFunctions.md):

* functions on function types (e.g. `AsTSV`)