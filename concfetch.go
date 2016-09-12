//!+

// Fetch prints the content found at each specified URL.
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

var semacount *int = flag.Int("sema", 20, "semaphore count")
var sema chan struct{}
var wg sync.WaitGroup

func main() {
	flag.Parse()
	sema = make(chan struct{}, *semacount)
	//fmt.Printf("sema=%d\n", *semacount)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		url := input.Text()
		wg.Add(1)
		go fetch(url)
	}
	wg.Wait()
}

//!+sema
// sema is a counting semaphore for limiting concurrency in fetch.

func fetch(url string) {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...
	//!-sema

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
