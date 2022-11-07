package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// reservoir, struct

func main() {
	br := bufio.NewReader(os.Stdin)
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(line) // reservoir.Add(..)
	}

}
