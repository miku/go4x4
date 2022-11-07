package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
)

type Todo struct {
	Id      int
	Title   string
	Done    bool
	Created time.Time
}

type Server struct {
	DB     *sqlx.DB
	Router mux.Router
}

func (s *Server) Routes() {
	s.Router.HandleFunc("/", s.handleIndex()).Method("GET")
	s.Router.HandleFunc("/new", s.handleNew()).Method("GET", "POST")
	s.Router.HandleFunc("/done/{id:[0-9]+}", s.handleDone()).Method("GET")
}

func (s *Server) handleIndex() http.HandlerFunc {
	// setup required for handler
	return func(w http.ResponseWriter, r *http.Request) {
		var todos []Todo
		err := s.DB.Select(&todos, "SELECT * FROM todo")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, v := range todos {
			log.Printf("%v", v)
		}
	}
}
func (s *Server) handleNew() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
func (s *Server) handleDone() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

type Config struct {
	HostPort string `default:"localhost:3000"`
	Database string `default:"todo.db"`
}

func main() {
	var config Config
	err := envconfig.Process("todo", &config)
	if err != nil {
		log.Fatal(err)
	}
	db, err := sqlx.Open("sqlite3", config.Database)
	if err != nil {
		log.Fatal(err)
	}
	srv := &Server{
		DB:     db,
		Router: mux.NewRouter(),
	}
	srv.Routes()
	log.Printf("starting service at http://%s", config.HostPort)
	log.Fatal(http.ListenAndServe(config.HostPort, srv))
}
