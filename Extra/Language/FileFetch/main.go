package main

import (
	"flag"
	"fmt"
)

// TODO:
// access ftp, credential, host, ...
// acesss file, print to stdout

// Fetch and print file to stdout.
func Fetch(...) error {
	
}

var (
	hostPort = flag.String("H", "ftp.ncbi.nlm.nih.gov:21", "ftp hostport")
)

func main() {
	// flag handling, args ...
	flag.Usage = func() {
		fmt.Printf("ftpcat prints a file from FTP to stdout\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Println(*hostPort)
}
