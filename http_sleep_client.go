package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

func main() {
	err, thread, tester := config()
	if err != nil {
		log.Fatal(err)
		return
	}
	wg := &sync.WaitGroup{}
	for i := 0; i < thread; i++ {
		wg.Add(1)
		go func() {
			tester.run()
			wg.Done()
		}()
	}
	wg.Wait()
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
	thread := flag.Int("thread", 10, "threads")
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
	}
	return nil, *thread, &t
}

func (t *tester) run() {
	for i := 0; t.loop <= 0 || i < t.loop; i++ {
		t.get()
	}
}

func (t *tester) get() {
	values := url.Values{}
	if t.min == t.max {
		values.Add("timer", strconv.Itoa(t.min))
	} else {
		values.Add("timer", strconv.Itoa(t.min+rand.Intn(t.max-t.min)))
	}

	req, err := http.NewRequest("GET", t.url+"/", nil)
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
