package main

import (
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
)

func BenchmarkServer(b *testing.B) {
	addr := "localhost:8000"
	url := "http://" + addr + "/greet"
	method := "GET"

	srv := newServer()
	srv.routes()
	go http.ListenAndServe(addr, srv.router)

	b.ResetTimer()
	wg := &sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(count int) {
			req, err := http.NewRequest(method, url, nil)
			if err != nil {
				b.Errorf("Creating request failed: %v", err)
			}

			client := new(http.Client)
			r, err := client.Do(req)
			if err != nil {
				b.Errorf("failed http.Client.Do: %v", err)
			}
			defer func() {
				err := r.Body.Close()
				if err != nil {
					b.Errorf("Close response.body failed: %v", err)
					return
				}
			}()

			if r.StatusCode != http.StatusOK {
				b.Errorf("Http request failed SatusCode: %v", r.StatusCode)
			}
			_, err = ioutil.ReadAll(r.Body)
			if err != nil {
				b.Errorf("Read response body failed : %v", err)
			}
			wg.Done()
			//log.Printf("[%v/%v]finished", count, b.N)
		}(i)
		wg.Wait()
	}
}
