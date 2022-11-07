package main

import (
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	msg := r.PostForm.Get("msg") // usr r.Form for query parameter
	io.WriteString(w, msg)
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
