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
	//separate each columns
	var out []string
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		line := scanner.Text()
		for i, v := range ci {
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
		out[i] = strings.TrimSpace(multispace_regexp.ReplaceAllString(v, " "))
	}
	return out
}

func getColValues(line string, ci [][]int, cn []string) []string {
	var out []string
	fmt.Println(line)
	for _, v := range ci {
		//fmt.Println(line)
		//fmt.Println(v)
		var word string
		if len(line) >= v[1] {
			word = line[v[0]:v[1]]
		} else if len(line) >= v[0] {
			word = line[v[0]:len(line)]
		}
		out = append(out, word)
	}
	//truncate spaces
	for i, v := range out {
		out[i] = strings.TrimSpace(multispace_regexp.ReplaceAllString(v, " "))
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
			var column_indicies [][]int
			var colnames []string
			for scanner.Scan() {
				//fmt.Println(scanner.Text())
				line := scanner.Text()
				//fmt.Println("[" + line + "]")
				if spaceline_regexp.MatchString(line) {
					//fmt.Println("set nil :" + line)
					header_lines = ""
					column_indicies = nil
				} else if isTableSeparator(line) {
					//fmt.Println("Talble Separator :" + line)
					//fmt.Println("Header Lines:")
					fmt.Print(header_lines)
					//fmt.Println("Table Separator:")
					fmt.Println(line)
					column_indicies = column_regexp.FindAllStringIndex(line, -1)
					//fmt.Println(column_indicies)
					colnames = getColNames(header_lines, column_indicies)
					//fmt.Println(colnames)
					//fmt.Println("=====")
				} else {
					//fmt.Println("else :" + line)
					if column_indicies == nil {
						header_lines += line + "\n"
					} else {
						values := getColValues(line, column_indicies, colnames)
						for i, v := range values {
							fmt.Println(colnames[i] + " : " + v)
						}
					}
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
