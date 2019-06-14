package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func showId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "My number is : %v\n", vars["id"])
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "handler1")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "handler2")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user/{id:[0-9]+}", showId)
	r.HandleFunc("/test/xxx", handler2)
	r.HandleFunc("/test/{id:x*}", handler1)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
