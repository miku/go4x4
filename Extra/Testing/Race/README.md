# Race detection

* [Typical Data Races](https://go.dev/doc/articles/race_detector#Typical_Data_Races)

```
$ go test -race .
==================
WARNING: DATA RACE
Write at 0x00c0000ae390 by goroutine 8:
  runtime.mapassign_faststr()
      /usr/local/go/src/runtime/map_faststr.go:203 +0x0
  github.com/miku/gows/Extra/Testing/Race.run.func1()
      /home/tir/code/miku/go4x4/Extra/Testing/Race/main.go:9 +0x50

Previous write at 0x00c0000ae390 by goroutine 7:
  runtime.mapassign_faststr()
      /usr/local/go/src/runtime/map_faststr.go:203 +0x0
  github.com/miku/gows/Extra/Testing/Race.run()
      /home/tir/code/miku/go4x4/Extra/Testing/Race/main.go:12 +0x132
  github.com/miku/gows/Extra/Testing/Race.TestRun()
      /home/tir/code/miku/go4x4/Extra/Testing/Race/main_test.go:6 +0x24
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1446 +0x216
  testing.(*T).Run.func1()
      /usr/local/go/src/testing/testing.go:1493 +0x47

Goroutine 8 (running) created at:
  github.com/miku/gows/Extra/Testing/Race.run()
      /home/tir/code/miku/go4x4/Extra/Testing/Race/main.go:8 +0x115
  github.com/miku/gows/Extra/Testing/Race.TestRun()
      /home/tir/code/miku/go4x4/Extra/Testing/Race/main_test.go:6 +0x24
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1446 +0x216
  testing.(*T).Run.func1()
      /usr/local/go/src/testing/testing.go:1493 +0x47

Goroutine 7 (running) created at:
  testing.(*T).Run()
      /usr/local/go/src/testing/testing.go:1493 +0x75d
  testing.runTests.func1()
      /usr/local/go/src/testing/testing.go:1846 +0x99
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1446 +0x216
  testing.runTests()
      /usr/local/go/src/testing/testing.go:1844 +0x7ec
  testing.(*M).Run()
      /usr/local/go/src/testing/testing.go:1726 +0xa84
  main.main()
      _testmain.go:47 +0x2e9
==================
2 b
1 a
--- FAIL: TestRun (0.00s)
    testing.go:1319: race detected during execution of test
FAIL
FAIL    github.com/miku/gows/Extra/Testing/Race 0.017s
FAIL
```