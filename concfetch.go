// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
	//"time"
)

var wg sync.WaitGroup

func main() {
	//start := time.Now()
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		url := input.Text()
		wg.Add(1)
		go fetch(url)
	}
	wg.Wait()
	//fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds())
}

//!+sema
// sema is a counting semaphore for limiting concurrency in fetch.
var sema = make(chan struct{}, 20)

func fetch(url string) {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...
	//!-sema

	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s %s\n", resp.Status, url)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
}

//!-
