//!+

// Kamehame sends http requests and prints results at each specified URL.
package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sync"
	"time"
)

// sema is a counting semaphore for limiting concurrency in http requesting.
var (
	sema chan struct{}
	wg   sync.WaitGroup
)

func wave(concurrency int, tps int, buf io.Reader) {
	var start = time.Now()
	sema = make(chan struct{}, concurrency)
	input := bufio.NewScanner(buf)

	for i := 0; input.Scan(); i++ {
		wg.Add(1)

		line := input.Text()
		rep := regexp.MustCompile(`[\s\t\r]+`)
		col := rep.Split(line, -1)

		var method = col[0]
		var url = col[1]

		/* for i := 0; i < len(col); i++ {
			fmt.Printf("col[%d]=%s\n", i, col[i])
		} */

		go fetch(method, url)

		// sleep difference expected time from elapsed time
		var wait = time.Duration(float64((i+1)/tps))*time.Second - time.Since(start)
		if wait > 0 {
			time.Sleep(wait)
		}
	}
	//fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	wg.Wait()
}

func fetch(method, url string) {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	defer wg.Done()
	start := time.Now()

	req, _ := http.NewRequest(method, url, nil)

	client := new(http.Client)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.Client.Do(): %v\n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	fmt.Printf("%d %s %6.6f\n", resp.StatusCode, url, time.Since(start).Seconds())
}

//!-
