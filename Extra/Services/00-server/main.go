package main

import (
	"io"
	"log"
	"net/http"
)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Hello", "World")
	// w.Write([]byte("Hello World!\n"))
	io.WriteString(w, "Hello World!\n")
	// io.Copy(w, strings.NewReader("Hello World!\n"))
}

func main() {

	log.Println("http://localhost:8000")
	http.ListenAndServe("localhost:8000", &handler{})

	// ListenAndServer returns an error, which should be checked.
}
