package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Fine."))
	case http.MethodPost:
		w.Write([]byte("Ok"))
	case http.MethodPut:
		w.Write([]byte("Updating."))
	case http.MethodDelete:
		w.Write([]byte("Sure?"))
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {

	log.Println("http://localhost:8000")

	// Convert the Echo function to a type that implements http.Handler
	h := http.HandlerFunc(handler)

	// Start a server listening on port 8000 and responding using Echo.
	if err := http.ListenAndServe("localhost:8000", h); err != nil {
		log.Fatalf("error: listening and serving: %s", err)
	}
}
