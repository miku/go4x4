# CSP

## Notes

* Communicating Sequential Processes (Hoare, 1978)
* Goroutines that communicate via channels
* Emphasis on communication, a goroutine does not address a target, but uses a channel
* *Do not communicate by sharing memory; instead, share memory by communicating.*

> One way to think about this model is to consider a typical single-threaded
> program running on one CPU. It has no need for synchronization primitives. Now
> run another such instance; it too needs no synchronization. Now let those two
> communicate; if the communication is the synchronizer, there's still no need
> for other synchronization. Unix pipelines, for example, fit this model
> perfectly. Although Go's approach to concurrency originates in Hoare's
> Communicating Sequential Processes (CSP), it can also be seen as a type-safe
> generalization of Unix pipes.

## Implemenatation in Go

* each physical core (M) is associated with a (virtual) processor (P)
* the program starts with one Goroutine (G)
* two run queues: global (GRQ), local (LRQ)

![](gpm.png)

The scheduler is part of the runtime and contained in any compiled binary.

* A Goroutine may be in one of three states: Waiting, Runnable or Executing.
* There are certain execution points, where scheduling decisions are made, e.g. on goroutine creation
* Performant, since context switching is happening at the application level -
  not on the OS level (context switch is)



## Code Review

* [NumCPU](example1/main.go)
