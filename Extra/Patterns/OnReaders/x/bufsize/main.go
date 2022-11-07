package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	b := bufio.NewReader(strings.NewReader("Hello World"))
	fmt.Println(b.Size())
}
