package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

func main() {
	REQ_URL := flag.String("url", "http://127.0.0.1:10182", "request url")
	thread := flag.Int("thread", 10, "threads")
	loop := flag.Int("loop", 0, "loop")
	min := flag.Int("min", 0, "min msec sleep")
	max := flag.Int("max", 100, "max msec sleep")
	keepalive := flag.Bool("keepalive", false, "keep alive Tcp connections")
	debug := flag.Bool("debug", false, "debug")
	flag.Parse()

	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 50
	http.DefaultClient.Timeout = 0
	//client := &http.Client{Timeout: time.Duration(10 * time.Second)}
	wg := &sync.WaitGroup{}
	t := &tester{
		client:    http.DefaultClient,
		REQ_URL:   *REQ_URL,
		loop:      *loop,
		min:       *min,
		max:       *max,
		keepalive: *keepalive,
		debug:     *debug,
	}
	for i := 0; i < *thread; i++ {
		wg.Add(1)
		go func() {
			t.do()
			wg.Done()
		}()
	}
	wg.Wait()
}

type tester struct {
	client    *http.Client
	REQ_URL   string
	loop      int
	min       int
	max       int
	keepalive bool
	debug     bool
}

func (t *tester) do() {
	if t.loop > 0 {
		for i := 0; i < t.loop; i++ {
			t.get()
		}
	} else {
		for {
			t.get()
		}
	}
}

func (t *tester) get() {
	values := url.Values{}
	values.Add("timer", strconv.Itoa(t.min+rand.Intn(t.max-t.min)))
	//fmt.Println(values.Encode())

	req, err := http.NewRequest("GET", t.REQ_URL+"/", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.URL.RawQuery = values.Encode()

	resp, err := t.client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if t.keepalive || t.debug {
		t.dump(resp)
		//fmt.Println("Finished!!")
	}
}

func (t *tester) dump(resp *http.Response) {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if t.debug {
		fmt.Print(string(b))
	}
}
