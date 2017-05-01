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

const interval float64 = 5 // 5 sec.

type connTable struct {
	start    map[string]float64
	duration map[string]float64
}

func newConnTable() *connTable {
	return &connTable{
		start:    make(map[string]float64),
		duration: make(map[string]float64),
	}
}

var token_regex = regexp.MustCompile(`[\s\t\r]+`)

func (ct *connTable) getLine(line string) {
	token := token_regex.Split(line, -1)
	timestamp, _ := strconv.ParseFloat(token[0], 64)
	src := token[2]
	flags := token[6]
	fmt.Printf("timestamp=%f src=%s flags=%s\n", timestamp, src, flags)
	switch {
	case strings.Contains(flags, "S"):
		fmt.Println("SYN")
		ct.start[src] = timestamp
	case strings.Contains(flags, "F"):
		fmt.Println("FIN")
		if _, ok := ct.start[src]; ok {
			ct.duration[src] = float64(time.Now().UnixNano()/1e9) - timestamp
		}
	}

}

func (ct *connTable) report(interval float64) {
	var total_time float64
	var total_count, start_count, finish_count int

	now := time.Now()
	unixNow := float64(now.UnixNano() / 1e9)

	total_count = len(ct.start)
	for k, start := range ct.start {
		if start > unixNow-interval {
			start_count += 1
		}
		if duration, ok := ct.duration[k]; ok { // Already Finished.
			total_time += duration
			finish_count += 1
			delete(ct.start, k)
			delete(ct.duration, k)
		} else { // Not finished yet.
			if start > unixNow-interval {
				total_time += unixNow - start
			} else {
				total_time += interval
			}
		}
	}

	fmt.Printf("%s\t%f\t%d\t%d\t%d\n",
		now.Format("2006/01/02 15:04:05.000 MST"),
		total_time, total_count, start_count, finish_count)
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

	tick := time.NewTicker(time.Second * 5).C //ToDo:support arg

	ct := newConnTable()

	for {
		select {
		case l := <-line:
			ct.getLine(l)
			//fmt.Println("Scanned one line.")
		case <-tick:
			ct.report(interval)
			fmt.Printf("[%s] Time has come!!\n", time.Now().Format("2006/01/02 15:04:05.000 MST"))
		}
	}
}
