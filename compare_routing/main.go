package main

import (
	"fmt"
	"net/http"
)

type server interface {
	routes()
	//handleHello() http.HandlerFunc
	//handleDefault() http.HandlerFunc
	handler() http.Handler
}

func main() {
	var s server
	s = newStandardServer()
	s.routes()
	http.ListenAndServe("localhost:8000", s.handler())
}

func handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	}
}

func handleDefault() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static", 301)
	}
}
