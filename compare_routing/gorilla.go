package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type gorillaServer struct {
	router *mux.Router
}

func newGorillaServer() *gorillaServer {
	return &gorillaServer{
		router: mux.NewRouter(),
	}
}

func (s *gorillaServer) handler() http.Handler {
	return s.router
}

func (s *gorillaServer) routes() {
	s.router.HandleFunc("/", handleDefault)
	s.router.HandleFunc("/greet", handleHello)
	//s.router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
}
