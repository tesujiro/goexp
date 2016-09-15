//!+

// Kamehame sends http requests and prints results at each specified URL.
package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// sema is a counting semaphore for limiting concurrency in http requesting.
var sema chan struct{}
var wg sync.WaitGroup

func wave(concurrency int, tps int, buf io.Reader) {
	var start = time.Now()
	sema = make(chan struct{}, concurrency)
	input := bufio.NewScanner(buf)

	for i := 0; input.Scan(); i++ {
		url := input.Text()
		wg.Add(1)
		go fetch(url)

		// sleep difference expected time from elapsed time
		var wait = time.Duration(float64((i+1)/tps))*time.Second - time.Since(start)
		if wait > 0 {
			time.Sleep(wait)
		}
	}
	//fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	wg.Wait()
}

func fetch(url string) {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	start := time.Now()
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%d %s %6.6f\n", resp.StatusCode, url, time.Since(start).Seconds())
	//fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds())
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
}

//!-
