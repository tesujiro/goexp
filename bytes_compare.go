package main

import (
	"fmt"
	"strings"
)

func sub(ch rune) bool {
	fmt.Printf("%v\n", ch)
	return false
}

func main() {
	s := "AbcDef"
	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", strings.IndexAny(s, "D"))
	fmt.Printf("%v\n", strings.IndexFunc(s, sub))
}
