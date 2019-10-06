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

type binarySlice []int

func (s *binarySlice) length() int {
	//var length = 0
	var bin int = 1
	var start int = 0
	for isLen(([]int)(*s), start) < 0 {
		start = start + bin
		bin *= 2
	}
	return s.binSearch(start-bin/2, bin/2)
}

// search length from start to start+bin-1
func (s *binarySlice) binSearch(start, bin int) int {
	//fmt.Printf("binSearch(%v, %v)\n", start, bin)
	if bin == 1 {
		return start
	}
	bin = bin / 2
	res := isLen(([]int)(*s), start+bin)
	switch {
	case res < 0:
		return s.binSearch(start+bin, bin)
	case res == 0:
		return start + bin
	default:
		return s.binSearch(start, bin)
	}
}
