package main

import "fmt"

func sub() func() int {
	i := 3
	f := func() int {
		return i * 100
	}
	i = 5
	return f
}

func main() {
	i := 3
	f := func() int {
		return i * 10
	}
	i = 5
	fmt.Printf("i=%v\n", i)
	fmt.Printf("f()=%v\n", f())
	fmt.Printf("sub()()=%v\n", sub()())
}
