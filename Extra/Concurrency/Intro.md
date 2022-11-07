# Concurrency Intro

## Communicating Sequential Processes

* Communicating Sequential Processes (Hoare, 1978)
* Share memory by communicating, don't communicate by sharing memory

This mantra contains "communication" - more message, less queue.

Core notions: goroutines, channels, select.

However, Go supports classic and CSP style concurrency primitives. When to use which?

* [https://github.com/miku/cignotes/blob/master/images/fig21.png](https://github.com/miku/cignotes/blob/master/images/fig21.png)
  (via [Concurrency in
  Go](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/))

## Runtime Scheduler

* the go runtime implements a scheduler, akin to an OS to manage goroutines
* there is always one goroutine running (i.e. main thread)
* M:N scheduler, utilizes multiple cores

> At any time, M goroutines need to be scheduled on N OS threads that runs on at
> most GOMAXPROCS numbers of processors. -- https://rakyll.org/scheduler/

* there are some guaruantee, but in general you can determine the scheduling
  order - not should you rely on it

Just as in an OS, goroutines will be in different states:

* running
* runnable
* blocked

Blocking can occur through:

* Sending and Receiving on Channel.
* Network I/O.
* Blocking System Call.
* Timers.
* Mutexes.

Previously, a variable `GOMAXPROCS` used to be 1, which would only use one OS
thread. It now defaults to the number of CPUs.

> The GOMAXPROCS variable limits the number of operating system threads that can
> execute user-level Go code simultaneously. -- https://pkg.go.dev/runtime

Looking into the scheduler with `GODEBUG`:

* [../../x/godebug/main.go](../../x/godebug/main.go)

## Goroutines

> A "go" statement starts the execution of a function call as an independent
> concurrent thread of control, or goroutine, within the same address space. 

> The function value and parameters are evaluated as usual in the calling
> goroutine, but unlike with a regular call, program execution does not wait for
> the invoked function to complete. Instead, the function begins executing
> independently in a new goroutine. When the function terminates, its goroutine
> also terminates. If the function has any return values, they are discarded
> when the function completes. 

* goroutine is a lightweight thread of execution
* a goroutine consumes about 2kb of memory
* you can start 1 million goroutines in about 2 seconds (and they'd consume 1/5
  of the RAM of an 16GB machine)

[embedmd]:# (../../x/manygoroutines/main.go)
```go
package main

import (
	"flag"
	"log"
	"runtime"
	"time"
)

var (
	N     = flag.Int("n", 10, "number of goroutines to start")
	sleep = flag.Duration("s", 100*time.Millisecond, "how long each goroutine sleeps")
)

func main() {
	flag.Parse()
	log.Printf("starting %d goroutines", *N)
	started := time.Now()
	for i := 0; i < *N; i++ {
		go f()
	}
	log.Printf("%d/%d goroutines started/running in %v", *N, runtime.NumGoroutine(), time.Since(started))
}

func f() {
	time.Sleep(*sleep)
}
```

* there is no outside control, you cannot "stop" a goroutine
* a goroutine will stop if the function it started ends
* a goroutine may run on one thread or many different threads (see: [LockOSThread](https://pkg.go.dev/runtime#LockOSThread))

## Misc

The channel axioms:

* A send to a nil channel blocks forever
* A receive from a nil channel blocks forever
* A send to a closed channel panics
* A receive from a closed channel returns the zero value immediately


## Considerations

* a goroutine is just a function, running in another thread, everything that is
  wrong in sequential code will be wrong in a goroutine as well
* e.g. an endless loop, the goroutine will never exit
* communicating errors is harder, we cannot just return an "error" since
  program execution continues immediately after the call
* there is no way to stop a goroutine from the outside

> **Cooperative multitasking**, also known as non-preemptive multitasking, is a
> style of computer multitasking in which the operating system never initiates
> a context switch from a running process to another process. Instead, in order
> to run multiple applications concurrently, processes voluntarily yield
> control periodically or when idle or logically blocked.

Package context was added only in Go 1.7, as a way to unify cancellation a bit.

* go concurrency gives you the parts, but no higher level constructs
* your library should only expose channels when it must
* keep concurrency under the hood but you API sequential
* callbacks are possible, but it is not the pattern of choice in go
* scheduling is done by the go runtime, indeterminate
* when multiple select cases are possible one is chosen at random

A [study](https://songlh.github.io/paper/go-study.pdf) found about the same
ratio of bugs in Go code using CSP vs classic structures like mutexes.


## Typical patterns

* fan out, fan out-fan it (workers and queue)
* bounded concurrency
* timeouts, [Timeout](timeout/main.go)
* signal handling, [Signal](signal/main.go)
