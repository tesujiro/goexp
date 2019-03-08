package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

func main() {
	mu1 := &sync.Mutex{}
	mu2 := &sync.Mutex{}

	// mutex1 -> mutex2
	go func() {
		fmt.Println("lock mu1")
		mu1.Lock()
		defer mu1.Unlock()
		fmt.Println("lock mu1 OK")

		time.Sleep(100 * time.Millisecond)

		fmt.Println("lock mu2")
		mu2.Lock()
		defer mu2.Unlock()
		fmt.Println("lock mu2 OK")
	}()

	// mutex2 -> mutex1
	go func() {
		fmt.Println("lock mu2")
		mu2.Lock()
		defer mu2.Unlock()
		fmt.Println("lock mu2 OK")

		time.Sleep(100 * time.Millisecond)

		fmt.Println("lock mu1")
		mu1.Lock()
		defer mu1.Unlock()
		fmt.Println("lock mu1 OK")
	}()

	time.Sleep(1 * time.Second)
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	os.Exit(130)
}
