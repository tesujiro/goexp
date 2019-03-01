package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestE2E(t *testing.T) {
	addr := "localhost:8000"

	srv := newServer()
	srv.routes()
	go http.ListenAndServe(addr, srv.router)

	cases := []struct {
		method string
		url    string
		status int
		body   string
		header map[string]string
	}{
		{method: "GET", url: "/greet", body: "Hello, World"},
		{method: "GET", url: "/static/", body: "hello, html!\n"},
		{method: "GET", url: "/static", status: http.StatusOK},            // http client does not return http.StatusMovedPermanently
		{method: "GET", url: "/static/index.html", status: http.StatusOK}, // http client does not return http.StatusMovedPermanently
		{method: "GET", url: "/no_page", status: http.StatusOK},           // http client does not return http.StatusMovedPermanently
	}
	for _, c := range cases {
		req, err := http.NewRequest(c.method, "http://"+addr+c.url, nil)
		if err != nil {
			t.Errorf("failed http.NewRequest %v", err)
		}

		client := new(http.Client)
		r, err := client.Do(req)
		if err != nil {
			t.Errorf("failed http.Client.Do %v", err)
		}
		defer r.Body.Close()

		//fmt.Printf("Result:%#v\n", r)
		if c.status == 0 && r.StatusCode != http.StatusOK ||
			c.status != 0 && r.StatusCode != c.status {
			fmt.Printf("result:%#v\n", r)
			t.Errorf("method:%v url:%v StatusCode: want:%v get:%v", c.method, c.url, c.status, r.StatusCode)
		}
		//fmt.Printf("header.Location:%#v\n", r.Header["Location"])
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("result:%#v\n", r)
			t.Errorf("method:%v url:%v Error by ioutil.ReadAll(). %v", c.method, c.url, err)
		}
		if c.body != "" && string(data) != c.body {
			fmt.Printf("result:%#v\n", r)
			t.Errorf("method:%v url:%v Data Error. [%v]", c.method, c.url, string(data))
		}
	}
}
