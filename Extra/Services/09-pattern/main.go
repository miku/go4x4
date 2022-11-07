package main

// These are just a few snippets. Read more under:
// https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831

// Shared dependencies are fields of the structure; Handlers hang off the server.
type server struct {
	db     *someDatabase
	router *someRouter
	email  EmailSender
}

func (s *server) handleSomething() http.HandlerFunc { ... }

// routes.go

func (s *server) routes() {
	s.router.HandleFunc("/api/", s.handleAPI())
	s.router.HandleFunc("/about", s.handleAbout())
	s.router.HandleFunc("/", s.handleIndex())

	s.router.HandleFunc("/admin", s.adminOnly(s.handleAdminIndex()))}


// Middleware are just Go functions

func (s *server) adminOnly(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if !currentUser(r).IsAdmin {
            http.NotFound(w, r)
            return
        }
        h(w, r)
    }
}

// request and response types specific to a function

func (s *server) handleSomething() http.HandlerFunc {
    type request struct {
        Name string
    }
    type response struct {
        Greeting string `json:"greeting"`
    }
    return func(w http.ResponseWriter, r *http.Request) {
        ...
    }
}