package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type httprouterServer struct {
	router *httprouter.Router
}

func newHttprouterServer() *httprouterServer {
	return &httprouterServer{
		router: httprouter.New(),
	}
}

func (s *httprouterServer) handler() http.Handler {
	return s.router
}

func (s *httprouterServer) handleDefault() func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		handleDefault(w, r)
	}
}

func (s *httprouterServer) handleHello() func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		handleHello(w, r)
	}
}

func (s *httprouterServer) routes() {
	s.router.Handle("GET", "/", s.handleDefault())
	s.router.Handle("GET", "/greet", s.handleHello())
	//s.router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
}
