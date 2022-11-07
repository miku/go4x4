# First Class Functions

Functions are first class objects. 

> Go supports first class functions, higher-order functions, user-defined
> function types, function literals, closures, and multiple return values. 

## Use case: Proxy Handling

* [https://pkg.go.dev/net/http#Transport](https://pkg.go.dev/net/http#Transport)
* [https://pkg.go.dev/net/http#RoundTripper](https://pkg.go.dev/net/http#RoundTripper)

The proxy field is function type, serving as an extension point in HTTP
processing regarding proxies.

> If two functions have the same signature (parameter and result types), they
> share one function type.

## Use case: HandlerFunc

Anything that looks like: `func(w http.ResponseWriter, r *http.Request)` is a
[HandlerFunc](https://pkg.go.dev/net/http#HandlerFunc).

> The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers.

## Use case: Mapper

Maps names to functions; example: a command line tool that supports a few
builtin tranformations.

Here, a `Mapper` is a function type describing a generic transformation (e.g.
from a blob of data to fields). Various implementations exist, the user can choose one.

```go

// Mapper maps a blob to an arbitrary number of fields, e.g. for (key,
// doc). We want fields, but we do not want to bake in TSV into each function.
type Mapper func([]byte) ([][]byte, error)

...

	availableMappers := map[string]skate.Mapper{
		// Add new mapper functions here. TODO: add more docs, and improve
		// composability, e.g. like middleware. Also improve naming.
		"id":   skate.Identity,
		"ff":   skate.CreateFixedMapper(*extraValue),
		"ti":   skate.MapperTitle,
		"tn":   skate.MapperTitleNormalized,
		"ty":   skate.MapperTitleNysiis,
		"ts":   skate.MapperTitleSandcrawler,
		"ur":   skate.MapperURLFromRef,
		"ru":   skate.MapperIdentURLFromRef,
		"cni":  skate.MapperContainerName,
		"cns":  skate.MapperContainerNameSandcrawler,
		"rcns": skate.MapperReleaseContainerName,
		"vcns": skate.MapperReleaseResolvedContainerName,
		"isbn": skate.MapperOpenLibraryReleaseNormalizedISBN,
		"cdxu": skate.MapperCdxSummary,
		"bref": skate.MapperBrefWork,
		"rewo": skate.MapperReleaseWork,
		"bidt": skate.MapperBrefIdentifierTable,
	}
```

Example for defining a function on a function:

* we would like to separate the concern of serializing the fields as TSV
* we also would like to have an adapter function with a slightly different signature

We can define a function on the function type, e.g. `AsTSV` - now every `Mapper`
can be serialized as TSV and we still have a function.

```go
// AsTSV serializes the result of a field mapper as TSV. This is a slim
// adapter, e.g. to parallel.Processor, which expects this function signature.
// A newline will be appended, if not there already.
//
// Anecdotally a parallelized implementation of a mapper can process around 300MiB/s.
func (f Mapper) AsTSV(p []byte) ([]byte, error) {
	var (
		fields [][]byte
		err    error
		b      []byte
	)
	if fields, err = f(p); err != nil {
		return nil, err
	}
	if len(fields) == 0 {
		return nil, nil
	}
	b = bytes.Join(fields, bTab)
	if len(b) > 0 && !bytes.HasSuffix(b, bNewline) {
		b = append(b, bNewline...)
	}
	return b, nil
}
```

For example we can use it in different library, that allows us to parallelize processing:

```
pp := parallel.NewProcessor(os.Stdin, os.Stdout, mapper.AsTSV)
```

The developer can focus on implementing a custom map function, serialization an
parallization are separated.