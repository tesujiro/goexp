package main

import (
	"net/http"
)

type standardServer struct {
	router *http.ServeMux
}

func newStandardServer() *standardServer {
	return &standardServer{
		router: http.NewServeMux(),
	}
}

func (s *standardServer) handler() http.Handler {
	return s.router
}

func (s *standardServer) routes() {
	s.router.HandleFunc("/", handleDefault)
	s.router.HandleFunc("/greet", handleHello)
	s.router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
}
