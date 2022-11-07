package main

import (
	"fmt"
	"log"
	"net/http"
)

// Exercise: Write a debug handler, that returns a response with information
// about the request. You can use the helper function: httputil.DumpRequest(r,
// false) - or return selected fields. Register the function with the DefaultServeMux.
//
// Return an HTTP 500, if something goes wrong.

// Debug prints out information about the request.

// Echo is a basic HTTP Handler.
func Echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You asked to %s %s\n", r.Method, r.URL.Path)
}

func main() {

	log.Println("http://localhost:8000")

	http.HandleFunc("/", Debug)

	// Start a server listening on port 8000 and responding using Echo.
	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		log.Fatalf("error: listening and serving: %s", err)
	}
}
