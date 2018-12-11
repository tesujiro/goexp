package main

import (
	"fmt"
	"net/http"
	"sync"
)

type server interface {
	routes()
	//handleHello() http.HandlerFunc
	//handleDefault() http.HandlerFunc
	handler() http.Handler
}

func startWebServer(s server, url string, wg *sync.WaitGroup) {
	wg.Add(1)
	s.routes()
	go func() {
		http.ListenAndServe(url, s.handler())
		wg.Done()
	}()
}

func main() {
	wg := &sync.WaitGroup{}

	// Standard Library Mux
	startWebServer(newStandardServer(), "localhost:8000", wg)

	// httprouter
	startWebServer(newHttprouterServer(), "localhost:8001", wg)

	// gorilla mux
	startWebServer(newGorillaServer(), "localhost:8002", wg)

	wg.Wait()
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func handleDefault(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/static", 301)
}
