package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func benchmarkStaticPage(s server, b *testing.B) {
	c := struct {
		method string
		url    string
		status int
		body   string
		header map[string]string
	}{method: "GET", url: "/greet", body: "Hello, World"}

	s.routes()
	req, err := http.NewRequest(c.method, c.url, nil)
	if err != nil {
		fmt.Printf("failed http.NewRequest %v", err)
		return
	}
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		s.handler().ServeHTTP(w, req)
		/*
			r := w.Result()
			//fmt.Printf("Result:%#v\n", r)
			if c.status == 0 && r.StatusCode != http.StatusOK ||
				c.status != 0 && r.StatusCode != c.status {
				fmt.Printf("result:%#v\n", r)
				t.Fatalf("method:%v url:%v StatusCode:%v", c.method, c.url, r.StatusCode)
			}
			//fmt.Printf("header.Location:%#v\n", r.Header["Location"])
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Printf("result:%#v\n", r)
				t.Fatalf("method:%v url:%v Error by ioutil.ReadAll(). %v", c.method, c.url, err)
			}
			if c.body != "" && string(data) != c.body {
				fmt.Printf("result:%#v\n", r)
				t.Fatalf("method:%v url:%v Data Error. [%v]", c.method, c.url, string(data))
			}
		*/
	}
}

func BenchmarkStaticPage_Standard(b *testing.B) {
	s := newStandardServer()
	benchmarkStaticPage(s, b)
}

func BenchmarkStaticPage_httprouter(b *testing.B) {
	s := newHttprouterServer()
	benchmarkStaticPage(s, b)
}
