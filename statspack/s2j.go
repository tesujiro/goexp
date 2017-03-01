package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"
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
			//out <- input.Text()
			out <- strings.Replace(input.Text(), "\r", "", -1)
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
					//fmt.Println("[" + s + "]")
				}
			}
		}
		out <- kv
		close(out)
	}()
	return out
}

var tab_regexp = regexp.MustCompile(`[^~ -]+|^[^~-]|^$`)
var column_regexp = regexp.MustCompile("[-]+|[~]+")
var spaceline_regexp = regexp.MustCompile("^[ \t]*$")
var multispace_regexp = regexp.MustCompile("[ \t]+")

func isTableSeparator(s string) bool {
	return !tab_regexp.MatchString(s)
}

func getColNames(s string, ci [][]int) []string {
	//merge each columns
	var out []string
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		line := scanner.Text()
		for i, v := range ci {
			//fmt.Println(line)
			//fmt.Println(v)
			var word string
			if len(line) >= v[1] {
				word = line[v[0]:v[1]]
			} else if len(line) >= v[0] {
				word = line[v[0]:len(line)]
			}
			//fmt.Println(" ==> Too Short")
			if len(out) <= i {
				out = append(out, word)
			} else {
				out[i] += " " + word
			}
		}
	}
	//truncate spaces
	for i, v := range out {
		out[i] = multispace_regexp.ReplaceAllString(v, " ")
	}
	return out
}

func parseSection(in <-chan KeyValue) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for kv := range in {
			fmt.Println("=====")
			fmt.Println(kv.Key)
			fmt.Println("=====")
			scanner := bufio.NewScanner(strings.NewReader(kv.Value))
			var header_lines string
			for scanner.Scan() {
				//fmt.Println(scanner.Text())
				line := scanner.Text()
				//fmt.Println("[" + line + "]")
				if spaceline_regexp.MatchString(line) {
					header_lines = ""
				} else if isTableSeparator(line) {
					fmt.Println("Table Separator:")
					fmt.Println(line)
					fmt.Println("Header Lines:")
					fmt.Println(header_lines)
					ci := column_regexp.FindAllStringIndex(line, -1)
					//fmt.Println(ci)
					fmt.Println(getColNames(header_lines, ci))
					fmt.Println("=====")
				} else {
					header_lines += line + "\n"
				}
			}
		}
	}()
	wg.Wait()
}

func main() {
	/*
		for kv := range separateSection(read(os.Stdin)) {
			fmt.Println("=====================")
			fmt.Println(kv.Key)
			fmt.Println("=====================")
			//fmt.Println(kv.Value)
		}
	*/
	//s := []string{"aaa", "bbb", "ccc"}
	//fmt.Println(s)
	parseSection(separateSection(read(os.Stdin)))
}
