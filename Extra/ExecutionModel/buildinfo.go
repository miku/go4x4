package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	info, _ := debug.ReadBuildInfo()
	fmt.Println(info)
}
