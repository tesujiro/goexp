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
	fmt.Printf("args: %#v\n", flag.Args())
	fmt.Printf("NArg:%v\n", flag.NArg())
	return fmt.Sprintf("bool:%v", *arg_bool)
}
