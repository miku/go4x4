package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Index is a basic HTTP Handler.
func Index(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{
		"Hello": "World",
		"Hi":    "Berlin",
	}
	w.Header().Set("Content-Type", "application/json")
	// b, err := json.Marshal(m)
	// if err != nil {
	// 	log.Println(err)
	// }
	// w.Write(b)
	json.NewEncoder(w).Encode(m)
	// fmt.Fprintf(w, "Welcome")
}

// Echo is a basic HTTP Handler.
func Echo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "You asked to %s %s\n", r.Method, r.URL.Path)
}

func main() {

	log.Println("http://localhost:8000")

	// Handle registers the handler for the given pattern in the DefaultServeMux.
	http.HandleFunc("/", Index)
	http.HandleFunc("/echo", Echo)

	// Start a server listening on port 8000 and responding using Echo.
	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		log.Fatalf("error: listening and serving: %s", err)
	}
}
