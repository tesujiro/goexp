package main

import "fmt"

var a [1024]int

func initArray() {
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
}

func sub(i int) *int {
	if i < 0 || i >= len(a) {
		return nil
	}
	return &a[i]
}

func print(i int) {
	fmt.Printf("i=%v\tsub(i)=%#v\t", i, sub(i))
	if sub(i) != nil {
		fmt.Printf("\t*sub(i)=%#v\n", *sub(i))
	} else {
		fmt.Printf("\n")
	}
}

func main() {
	initArray()
	print(-1)
	print(100)
	print(2000)
	*sub(100) = 200
	print(100)
}
