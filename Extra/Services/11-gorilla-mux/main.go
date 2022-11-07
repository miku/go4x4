package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Echo is a basic HTTP Handler.
func Echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You asked to %s %s\n", r.Method, r.URL.Path)
}

func main() {

	log.Println("http://localhost:8000")

	// More flexible mux: https://github.com/gorilla/mux#matching-routes
	mux := mux.NewRouter()

	mux.HandleFunc("/index", Echo)

	// Start a server listening on port 8000 and responding using Echo.
	if err := http.ListenAndServe("localhost:8000", mux); err != nil {
		log.Fatalf("error: listening and serving: %s", err)
	}
}
