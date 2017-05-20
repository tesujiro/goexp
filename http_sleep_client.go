package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func main() {
	err, threads, tester := config()
	if err != nil {
		log.Fatal(err)
		return
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	//ctx, timeout := context.WithTimeout(ctx, 3*time.Second)
	//defer timeout()
	wg := &sync.WaitGroup{}

	// tester goroutin start
	got := make(chan struct{}, threads)
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			tester.run(ctx, got)
			wg.Done()
		}()
	}

	// wait goroutine start
	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	// select all the events
	var counter int
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
L:
	for {
		select {
		case <-got:
			counter += 1
		case <-done:
			break L
		case <-ctx.Done():
			//fmt.Printf("Context done:\n")
			break L
		case s := <-sig:
			fmt.Printf("Got signal:%d\n", s)
			cancel()
		}
	}
	fmt.Printf("%d Requests returned OK.\n", counter)

}

type tester struct {
	client    *http.Client
	url       string
	loop      int
	min       int
	max       int
	keepalive bool
	debug     bool
}

func config() (error, int, *tester) {
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 50
	http.DefaultClient.Timeout = 0
	//client := &http.Client{Timeout: time.Duration(10 * time.Second)}
	t := tester{
		client: http.DefaultClient,
	}
	threads := flag.Int("thread", 10, "threads")
	flag.StringVar(&t.url, "url", "http://127.0.0.1:80", "request url")
	flag.IntVar(&t.loop, "loop", 0, "loop")
	flag.IntVar(&t.min, "min", 0, "min msec sleep")
	flag.IntVar(&t.max, "max", 100, "max msec sleep")
	flag.BoolVar(&t.keepalive, "keepalive", false, "keep alive Tcp connections")
	flag.BoolVar(&t.debug, "debug", false, "debug")
	flag.Parse()
	if t.min > t.max {
		err := fmt.Errorf("Error: min > max")
		return err, 0, nil
	} else if *threads < 0 || t.min < 0 || t.max < 0 || t.loop < 0 {
		err := fmt.Errorf("Error: negative number")
		return err, 0, nil
	}
	return nil, *threads, &t
}

func (t *tester) run(ctx context.Context, got chan<- struct{}) {
	for i := 0; t.loop <= 0 || i < t.loop; i++ {
		if err := t.get(); err == nil {
			got <- struct{}{}
		} else {
			log.Println(err)
		}
	}
}

func (t *tester) get() error {
	// set Query parameter "timer=n"
	values := url.Values{}
	if t.min == t.max {
		values.Add("timer", strconv.Itoa(t.min))
	} else {
		values.Add("timer", strconv.Itoa(t.min+rand.Intn(t.max-t.min)))
	}

	req, err := http.NewRequest("GET", t.url+"/", nil)
	if err != nil {
		return err
	}
	req.URL.RawQuery = values.Encode()

	resp, err := t.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if t.keepalive || t.debug {
		t.dump(resp)
	}
	return nil
}

func (t *tester) dump(resp *http.Response) {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	if t.debug {
		log.Print(string(b))
	}
}
