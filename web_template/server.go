package main

import (
	"fmt"
	"net/http"
)

type server struct {
	router *http.ServeMux
}

func newServer() *server {
	return &server{
		router: http.NewServeMux(),
	}
}

func main() {
	s := newServer()
	s.routes()
	http.ListenAndServe("localhost:8000", s.router)
}

func (s *server) routes() {
	s.router.HandleFunc("/greet", s.handleHello())
}

func (s *server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	}
}
