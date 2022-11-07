## Goroutines

Goroutines are functions that are created and scheduled to be run independently by the Go scheduler. The Go scheduler is responsible for the management and execution of goroutines.

## Notes

* Goroutines are functions that are scheduled to run independently.
* We must always maintain an account of running goroutines and shutdown cleanly. You can't kill a goroutine from outside.
* Concurrency is not parallelism.
	* Concurrency is about dealing with lots of things at once.
	* Parallelism is about doing lots of things at once.

_"Parallelism is TRYING to do two things at the same time. Concurrency is arranging it so you can do two things at the same time." - Dave Cheney_

## Design Guidelines

### Concurrent Software Design

Concurrency means “out of order” execution. Taking a set of instructions that
would otherwise be executed in sequence and finding a way to execute them out of
order and still produce the same result. For the problem in front of you, it has
to be obvious that out of order execution would add value. When I say value, I
mean add enough of a performance gain for the complexity cost. Depending on your
problem, out of order execution may not be possible or even make sense.

It’s also important to understand that [concurrency is not the same as
parallelism](https://blog.golang.org/concurrency-is-not-parallelism).

Parallelism means executing two or more instructions at the same time. This is a
different concept from concurrency. Parallelism is only possible when you have
at least 2 operating system (OS) and hardware threads available to you and you
have at least 2 Goroutines, each executing instructions independently on each
OS/hardware thread.

Both you and the runtime have a responsibility in managing the concurrency of
the application. You are responsible for managing these three things when
writing concurrent software:

**Design Philosophy:**

* Keep integrity
* Monitor critical points
* Rate limit, set general limits
* Use timeouts

Index of the three part series:

* 1) [Scheduling In Go : Part I - OS Scheduler](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html)
* 2) [Scheduling In Go : Part II - Go Scheduler](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html)
* 3) [Scheduling In Go : Part III - Concurrency](https://www.ardanlabs.com/blog/2018/12/scheduling-in-go-part3.html)

## Diagrams

### How the scheduler works.

![Ardan Labs](scheduler.png?v=2)

## Links

* [Scheduling In Go - Part I](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html) - William Kennedy
* [Scheduling In Go - Part II](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html) - William Kennedy
* [Advanced Go Concurrency Patterns](https://blog.golang.org/advanced-go-concurrency-patterns) - Sameer Ajmani
* [Go Concurrency Patterns: Context](https://blog.golang.org/context) - Sameer Ajmani
* [Concurrency is not parallelism](https://blog.golang.org/concurrency-is-not-parallelism) - Rob Pike
* [Go, for Distributed Systems](https://talks.golang.org/2013/distsys.slide) - Russ Cox
* [Go 1.5 GOMAXPROCS Default](https://docs.google.com/document/d/1At2Ls5_fhJQ59kDK2DFVhFu3g5mATSXqqV5QrxinasI/edit)
* [Concurrency, Goroutines and GOMAXPROCS](https://www.ardanlabs.com/blog/2014/01/concurrency-goroutines-and-gomaxprocs.html) - William Kennedy
* [The Linux Scheduler: a Decade of Wasted Cores](http://www.ece.ubc.ca/~sasha/papers/eurosys16-final29.pdf)
* [Explanation of the Scheduler](https://news.ycombinator.com/item?id=12460807)
* [15 Years of Concurrency](http://joeduffyblog.com/2016/11/30/15-years-of-concurrency/) - Joe Duffy
* [How does the golang scheduler work?](https://www.quora.com/How-does-the-golang-scheduler-work/answer/Ian-Lance-Taylor) - Ian Lance Taylor
* [The Scheduler Saga](https://www.youtube.com/watch?v=YHRO5WQGh0k) - Kavya Joshi

## Code Review

* [Goroutines and concurrency](example1/example1.go) ([Go Playground](https://play.golang.org/p/4n6G3uRDc83))
* [Goroutine time slicing](example2/example2.go) ([Go Playground](https://play.golang.org/p/QtNVo1nb4uQ))
* [Goroutines and parallelism](example3/example3.go) ([Go Playground](https://play.golang.org/p/ybZ84UcLW81))

## Exercises

### Exercise 1

**Part A** Create a program that declares two anonymous functions. One that counts down from 100 to 0 and one that counts up from 0 to 100. Display each number with an unique identifier for each goroutine. Then create goroutines from these functions and don't let main return until the goroutines complete.

**Part B** Run the program in parallel.

* [Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/O0FB2gd6-7d)) |
* [Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/uZlHjwf2CXY))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
