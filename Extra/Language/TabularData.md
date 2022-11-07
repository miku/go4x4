# Tabular Data

Go has support for classic CSV and additionally some libraries inspired by data
frames in other languages (e.g. R, Python, ...).

* [csv](https://pkg.go.dev/encoding/csv)
* https://github.com/rocketlaunchr/dataframe-go
* https://github.com/go-gota/gota
* https://github.com/tobgu/qframe

Depending on the size and use case, you can load you data into memory at once or
you may need to use a streaming approach.

## CSV

Works with any reader. Two main ways to read:

* [Read](https://pkg.go.dev/encoding/csv#Reader.Read)
* [ReadAll](https://pkg.go.dev/encoding/csv#Reader.ReadAll)

Analogously, you can write a single record or all records at once.

Internally, CSV writes will be buffered - so you'll need to call
[Flush](https://pkg.go.dev/encoding/csv#Writer.Flush) at the end (easy to
forget).

```go
// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		Comma: ',',
		w:     bufio.NewWriter(w),
	}
}

```

## Custom Struct Tags

There are libraries for bringing JSON and XML like struct tag support to CSV.

* [csvutil](https://nicedoc.io/jszwec/csvutil)

Examples:

* [TabularData/CsvUtilMarshalNested](TabularData/CsvUtilMarshalNested/main.go)
* [TabularData/CsvUtilUnmarshal](TabularData/CsvUtilUnmarshal/main.go)
* [TabularData/CsvUtilMap](TabularData/CsvUtilMap/main.go)


## How to write your custom struct tag support

You can write CSV support yourself, e.g. by using a helper library to access
structs.

Example:

* [TabularData/Decoder](TabularData/Decoder/main.go)