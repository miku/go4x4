package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

var messages []string

func postHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("received message from %v", r.RemoteAddr)
	r.ParseForm()
	msg := r.PostForm.Get("msg")
	if msg == "" {
		b, _ := httputil.DumpRequest(r, false)
		log.Printf("not our request: %s", string(b))
	}
	messages = append(messages, msg)
	log.Printf("now %v messages", len(messages))

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if len(messages) == 0 {
		io.WriteString(w, `no messages yet: use curl -XPOST 10.1.0.37/p -d 'msg=Hello' to post`)
	}
	for _, m := range messages {
		if _, err := io.WriteString(w, m+"\n"); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func main() {

	log.Println(`curl -XPOST http://10.1.0.37:8000/p -d 'msg=Hello'`)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/p", postHandler)

	http.ListenAndServe("0.0.0.0:8000", nil)
}
