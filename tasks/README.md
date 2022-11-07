# Tasks

Some sample tasks.

## Hello World

A hello world.

## Variable Declarations

Go has the "var" keyword, which can be used to declare variables, as well as a shorthand notation.
Use 4 different styles of variable declaration and write a program that prints out:

```
0 b 123 d
```

## Random Image

Generate a random PNG image, similar to the following examples.

![](../static/randomimg-1.png)

![](../static/randomimg-2.png)


## Word frequencies

Create a program that implements a word counter. Use a map to map tokens to their count.

```shell
$ links -dump https://go.dev/ref/spec | go run tasks/wordfreq/main.go | sort -nr | head -15
1597    the
997     a
902     of
889     type
832     is
568     and
464     //
446     to
400     =
346     in
319     or
317     an
316     be
310     are
285     The
...
```

## Switch 

Switch can be used with variables of many types, or no variable at all. Write
a short program, that get current day and depending on the hour print "Good
morning" (before noon), "Good afternoon" (before 5pm) or "Good evening".

## Bitcoin price

Write a program that fetches the current bitcoin price from an API and prints it out.

```
$ go run btcprice.go 

```

You could use [JSONGen](https://github.com/bemasher/JSONGen) to generate a struct for some example JSON data. 