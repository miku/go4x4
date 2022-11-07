package main

import (
	"log"
	"net/http"
)

func main() {

	log.Println("http://localhost:8000")
	http.ListenAndServe("localhost:8000",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello World!\n"))
		}),
	)

	// ListenAndServer returns an error, which should be checked.
}
