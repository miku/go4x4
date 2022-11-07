package main

import (
	"fmt"
	"log"
	"net/http"
)

// Service may hold references to shared objects, e.g. database handle.
type Service struct {
	Name string
}

// Index is a basic HTTP Handler.
func (s *Service) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to %s", s.Name)
}

// Echo is a basic HTTP Handler.
func (s *Service) Echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You asked to %s %s\n", r.Method, r.URL.Path)
}

func main() {

	log.Println("http://localhost:8000")

	svc := Service{Name: "Service 1.0"}

	// Handle registers the handler for the given pattern in the DefaultServeMux.
	http.HandleFunc("/", svc.Index)
	http.HandleFunc("/echo", svc.Echo)

	// Start a server listening on port 8000 and responding using Echo.
	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		log.Fatalf("error: listening and serving: %s", err)
	}
}
