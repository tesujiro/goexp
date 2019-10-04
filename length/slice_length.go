package main

import "fmt"

type lengther interface {
	length() int
}

func main() {
	a := make(naiveSlice, 10)
	fmt.Println(a.length())
}

func isLen(s []int, length int) int {
	//return len(s) - length
	l := len(s)
	if length > l {
		return 1
	} else if length < l {
		return -1
	}
	return 0
}

type naiveSlice []int

func (s *naiveSlice) length() int {
	var length int
	for length = 0; isLen(([]int)(*s), length) < 0; length++ {
	}
	return length
}
