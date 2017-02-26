package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

type Section struct {
	Name string
	Body string
}

type KeyValue struct {
	Key   string
	Value string
}

type StatspackReport struct {
}

//format_json := ` `
const SECTION_SEPARATOR = "                                    -------------------------------------------------------------"

func read(in io.Reader) <-chan string {
	out := make(chan string)
	go func() {
		input := bufio.NewScanner(in)
		for input.Scan() {
			out <- input.Text()
			//out <- strings.Replace(input.Text(), "\r", "", -1)
		}
		close(out)
	}()
	return out
}

func isEmptyLine(s string) {
}

var ss_regexp = regexp.MustCompile("([^~]+|^$)")

func isSubSectionSeparator(s string) bool {
	return !ss_regexp.MatchString(s)
	//return ss_regexp.MatchString(s)
}

var tab_regexp = regexp.MustCompile("[^ -]+")

func isTableSeparator(s string) bool {
	return !tab_regexp.MatchString(s) && s != ""
}

func separateSection(in <-chan string) <-chan KeyValue {
	out := make(chan KeyValue)
	go func() {
		var kv KeyValue
		re := regexp.MustCompile("[ \r\f]")
		for s := range in {
			if s == SECTION_SEPARATOR {
				out <- kv
				kv = KeyValue{}
			} else {
				if kv.Key == "" && re.ReplaceAllString(s, "") != "" {
					//fmt.Println("PATH!![" + s + "]")
					kv.Key = s
				}
				kv.Value += s + "\n" //ToDo: performance
				if isSubSectionSeparator(s) {
					fmt.Println("[" + s + "]")
				}
			}
		}
		out <- kv
		close(out)
	}()
	return out
}

func main() {
	/*
		for s := range read(os.Stdin) {
			fmt.Println(s)
		}
	*/
	for kv := range separateSection(read(os.Stdin)) {
		fmt.Println("=====================")
		fmt.Println(kv.Key)
		fmt.Println("=====================")
		//fmt.Println(kv.Value)
	}
}
