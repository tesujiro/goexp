package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "&a&b\\&c"
	fmt.Println(str)
	//re := regexp.MustCompile(`([!\\])\&`)
	re := regexp.MustCompile(`(^|[^\\])\&`)
	//re := regexp.MustCompile(`([^a])\&`)
	//re := regexp.MustCompile(`(!a)\&`)
	//re := regexp.MustCompile(`(a)\&`)
	//re := regexp.MustCompile(`\&`)
	//re := regexp.MustCompile(`^[\\\\]\&`)
	//re := regexp.MustCompile(`^\&`)
	//re := regexp.MustCompile(`(^[\\\\])&`)
	//re := regexp.MustCompile(`(^(\\\\))&`)
	//re := regexp.MustCompile(`(^\\\\)&`)
	//re := regexp.MustCompile(`\\&`)
	fmt.Println(re.ReplaceAllString(str, "${1}X"))
	//fmt.Println(re.ReplaceAllString(str, "X"))
}
