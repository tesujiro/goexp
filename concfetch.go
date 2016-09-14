//!+

// Concurrent Fetch sends http request and prints results at each specified URL.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

// sema is a counting semaphore for limiting concurrency in fetch.
var semacount *int = flag.Int("sema", 1, "semaphore count")
var sema chan struct{}

var tps *int = flag.Int("tps", 60, "target transaction / second")
var wg sync.WaitGroup

func main() {
	flag.Parse()
	sema = make(chan struct{}, *semacount)
	var interval = time.Duration(float64(1000)/float64(*tps)*0.9) * time.Millisecond
	//fmt.Printf("interval=%s (%f)\n", interval, interval)
	//fmt.Printf("sema=%d\n", *semacount)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		start := time.Now()
		url := input.Text()
		wg.Add(1)
		go fetch(url)
		//time.Sleep(interval)
		time.Sleep(interval - time.Since(start))
		//start = time.Now()
	}
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
