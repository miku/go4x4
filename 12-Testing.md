# Testing

* the go tool and standard library provide a basic testing support, no extra library needed to get started
* idea: let test code look like normal code; more verbose, but uniform

## Conventions

* put a file alongside code with suffix: `_test.go`, e.g.

```
sum.go
sum_test.go
```

There is a "testing" package and a type `testing.T` that is passed into the test function.

## Running tests




## Mocking

* test doubles for external dependencies

# Benchmarking and Profiling

## Benchmarks

Benchmarks can be written alongside tests.

* [Benchmarks](https://golang.org/pkg/testing/#hdr-Benchmarks)

Example benchmark:

```go
func BenchmarkHello(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("hello")
    }
}
```

> The benchmark function must run the target code b.N times. During benchmark
> execution, b.N is adjusted until the benchmark function lasts long enough to
> be timed reliably. The output

```
BenchmarkHello    10000000    282 ns/op
```

means that the loop ran 10000000 times at a speed of 282 ns per loop.

## Offline Profiling

Package [runtime/pprof](https://golang.org/pkg/runtime/pprof/) in the standard
library allows to collect profiling data.

It can run alongside tests:

```
$ go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
```

Or in standalong programs as well, by using flags (typically called `cpuprofile`
and `memprofile`, respectively):

```go
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
    flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal("could not create CPU profile: ", err)
        }
        defer f.Close()
        if err := pprof.StartCPUProfile(f); err != nil {
            log.Fatal("could not start CPU profile: ", err)
        }
        defer pprof.StopCPUProfile()
    }

    // ... rest of the program ...

    if *memprofile != "" {
        f, err := os.Create(*memprofile)
        if err != nil {
            log.Fatal("could not create memory profile: ", err)
        }
        defer f.Close()
        runtime.GC() // get up-to-date statistics
        if err := pprof.WriteHeapProfile(f); err != nil {
            log.Fatal("could not write memory profile: ", err)
        }
    }
}
```

Here's an example of a CPU profile output, rendered as PNG.

* [span-import.png](https://raw.githubusercontent.com/miku/span/master/docs/span-import.0.1.253.png)

![](span-import.0.1.253.png)

## Live Profiling

### Exposing handlers


* Run webserver for getting Go profiles (with `import _ "net/http/pprof"`)
* Fetch profile from localhost:$PORT/debug/pprof/$PROFILE_TYPE
* Use `go tool pprof` to analyze profile

### Exporing variables

A standardized interface to expose data over HTTP.

* [pkg/expvar](https://golang.org/pkg/expvar/)

> Package expvar provides a standardized interface to public variables, such as
> operation counters in servers. It exposes these variables via HTTP at
> /debug/vars in JSON format.

You can register these variables for custom metrics, then request them via HTTP.

There is a command line tool, called
[expvarmon](https://github.com/divan/expvarmon) for a basic visualization.

The `debug/vars` endpoint is fixed:

```go
// https://golang.org/src/expvar/expvar.go
func init() {
	http.HandleFunc("/debug/vars", expvarHandler)
	Publish("cmdline", Func(cmdline))
	Publish("memstats", Func(memstats))
}
```
