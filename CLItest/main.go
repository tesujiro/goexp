package main

import (
	"flag"
	"fmt"
)

var arg_bool = flag.Bool("b", false, "bool arg")

func main() {
	parse()
}

func parse() string {
	flag.Parse()
	return fmt.Sprintf("bool:%v", *arg_bool)
}
