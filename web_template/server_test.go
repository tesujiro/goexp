package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreet(t *testing.T) {
	srv := newServer()
	srv.routes()
	req, err := http.NewRequest("GET", "/greet", nil)
	if err != nil {
		t.Fatalf("failed http.NewRequest %v", err)
	}
	w := httptest.NewRecorder()
	srv.router.ServeHTTP(w, req)
	r := w.Result()
	//fmt.Printf("Result:%#v\n", r)
	if r.StatusCode != http.StatusOK {
		t.Fatalf("StatusCode:%v", r.StatusCode)
	}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("Error by ioutil.ReadAll(). %v", err)
	}
	if "Hello, World" != string(data) {
		t.Fatalf("Data Error. %v", string(data))
	}
}
