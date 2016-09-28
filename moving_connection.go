package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	//"strings"
	"strconv"
	"time"
)

const (
	time_fmt = "02/Jan/2006:15:04:05 -0700"
)

var (
	start                *string = flag.String("start", time.Now().Format(time_fmt), "start time")
	end                  *string = flag.String("end", time.Now().Format(time_fmt), "end time")
	unit_ms              *int    = flag.Int("unit", 1000, "unit millisecond")
	start_time, end_time time.Time
	unit                 time.Duration
	conn_count           = make(map[time.Time]int)
)

var rep = regexp.MustCompile(`\.[0-9]+`)

func s2t(str string) time.Time {
	str_t := rep.ReplaceAllString(str, "")
	t, err := time.Parse(time_fmt, str_t)
	if err != nil {
		log.Fatalf("s2t error %s", err)
	}
	under_sec := rep.FindString(str)
	if under_sec != "" {
		f64, err := strconv.ParseFloat(under_sec, 64)
		if err != nil {
			log.Fatalf("strconv.ParseFloat error %s", err)
		}
		t = t.Add(time.Millisecond * time.Duration(f64*1000))
	}
	return t
}

func begin() {
	flag.Parse()
	start_time = s2t(*start)
	end_time = s2t(*end)
	unit = time.Millisecond * time.Duration(*unit_ms)
	fmt.Printf("start_time=%v end_time=%v unit=%v\n", start_time, end_time, unit)
	// init connction counter
	//for t := start_time; t.Before(end_time) || t.Equal(end_time); t = t.Add(unit) {
	for t := start_time; !t.After(end_time); t = t.Add(unit) {
		conn_count[t] = 0
		//fmt.Println(t)
	}
}

func print_stat() {
	for t := start_time; t.Before(end_time) || t.Equal(end_time); t = t.Add(unit) {
		fmt.Printf("%v\t%d\n", t, conn_count[t])
	}
}

var rep_count = regexp.MustCompile(`[\s\t\r]+`)

func count(in io.Reader) {
	input := bufio.NewScanner(in)
	for input.Scan() {
		line := input.Text()
		//fmt.Printf("line=%s\n", line)
		result := rep_count.Split(line, -1)
		trn_start := s2t(fmt.Sprintf("%s %s", result[3][1:], result[4][:5]))
		//fmt.Printf("trn_start=%v\n", trn_start)
		f32, err := strconv.ParseFloat(result[len(result)-1], 32)
		//f64, err := strconv.ParseFloat(result[len(result)-1], 64)
		if err != nil {
			log.Fatalf("strconv.ParseFloat error %s", err)
		}
		response := time.Millisecond * time.Duration(f32*1000)
		//response := time.Millisecond * time.Duration(f64*1000)
		trn_end := trn_start.Add(response)
		//fmt.Printf("response=%v\n", response)
		for t := start_time; !t.After(end_time); t = t.Add(unit) {
			if !t.Before(trn_start) && !t.After(trn_end) {
				conn_count[t] += 1
			}
		}
	}
}

func main() {
	begin()
	count(os.Stdin)
	print_stat()
}
