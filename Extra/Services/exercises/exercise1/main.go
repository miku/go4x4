package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

// Debug show the request.
func Debug(w http.ResponseWriter, r *http.Request) {
	b, err := httputil.DumpRequest(r, false)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Write(b)
}

// Echo is a basic HTTP Handler.
func Echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You asked to %s %s\n", r.Method, r.URL.Path)
}

func main() {

	log.Println("http://localhost:8000")

	// Handle registers the handler for the given pattern in the DefaultServeMux.
	http.HandleFunc("/", Debug)
	http.HandleFunc("/echo", Echo)

	// Start a server listening on port 8000 and responding using Echo.
	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		log.Fatalf("error: listening and serving: %s", err)
	}
}
