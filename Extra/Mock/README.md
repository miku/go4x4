# Mocking

Example: [https://github.com/golang/mock](https://github.com/golang/mock)

```
$ go install github.com/golang/mock/mockgen@v1.6.0
```

Two modes:

* source (source file)
* reflect (an import path, and a comma-separated list of symbols.)

> Source mode generates mock interfaces from a source file. It is enabled by
> using the -source flag. 

```
$ mockgen \
  --source=foo.go \
  --destination=mock_test.go \
  --package=async_test
```

> Hamcrest?