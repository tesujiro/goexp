package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const interval float64 = 1 // 5 sec.
//const interval float64 = 5 // 5 sec.

type connTable struct {
	start         map[string]float64
	finish        map[string]float64
	over_interval map[string]bool
}

func newConnTable() *connTable {
	return &connTable{
		start:         make(map[string]float64),
		finish:        make(map[string]float64),
		over_interval: make(map[string]bool),
	}
}

var token_regex = regexp.MustCompile(`[\s\t\r]+`)

func (ct *connTable) getLine(line string) {
	token := token_regex.Split(line, -1)
	timestamp, _ := strconv.ParseFloat(token[0], 64)
	src := token[4]
	flags := token[6]
	//fmt.Printf("timestamp=%f src=%s flags=%s\n", timestamp, src, flags)
	switch {
	case strings.Contains(flags, "S"):
		//fmt.Println("SYN")
		ct.start[src] = timestamp
	case strings.Contains(flags, "F"):
		//fmt.Println("FIN")
		if start, ok := ct.start[src]; ok {
			ct.finish[src] = timestamp - start
			delete(ct.start, src)
		}
	}
}

func (ct *connTable) report() {
	var total_time float64
	var total_count, start_count, finish_count, current_conn int

	now := time.Now()
	unixNow := float64(now.UnixNano()) / float64(1e9)

	total_count = len(ct.start) + len(ct.finish)
	start_count = total_count - len(ct.over_interval)
	finish_count = len(ct.finish)
	current_conn = len(ct.start)
	for k, start := range ct.start {
		total_time += unixNow - start
		ct.over_interval[k] = true
	}
	for k, duration := range ct.finish {
		total_time += duration
		delete(ct.finish, k)
		if _, ok := ct.over_interval[k]; ok {
			delete(ct.over_interval, k)
		}
	}

	fmt.Printf("[%s]\ttotal=%f\tcount=%d\tstart=%d\tfinish=%d\tcurrent=%d\n",
		now.Format("2006/01/02 15:04:05.000 MST"),
		total_time, total_count, start_count, finish_count, current_conn)
}

func readLine(in io.Reader, line chan string) {
	input := bufio.NewScanner(in)
	for input.Scan() {
		line <- input.Text()
	}
}

func main() {
	line := make(chan string, 1)
	go readLine(os.Stdin, line)

	tick := time.NewTicker(time.Second * time.Duration(interval)).C //ToDo:support arg

	ct := newConnTable()

	for {
		select {
		case l := <-line:
			ct.getLine(l)
			//fmt.Println("Scanned one line.")
		case <-tick:
			ct.report()
			//fmt.Printf("[%s] Time has come!!\n", time.Now().Format("2006/01/02 15:04:05.000 MST"))
		}
	}
}
