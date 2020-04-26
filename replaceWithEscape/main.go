package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "&a\\&"
	fmt.Println(str)
	re := regexp.MustCompile(`^\&`)
	//re := regexp.MustCompile(`^&`)
	fmt.Println(re.ReplaceAllString(str, "X"))
}
