package main

type lengther interface {
	length() int
}

type base struct {
	size int
}

func (b base) shorterThan(length int) int {
	if b.size < length {
		return 1
	} else if b.size > length {
		return -1
	}
	return 0
}

type naiveList base

func newNaiveList(l int) naiveList {
	return naiveList{size: l}
}

func (n naiveList) length() int {
	var length int
	for length = 0; base(n).shorterThan(length) < 0; length++ {
	}
	return length
}

type binSearchList base

func newBinSearchList(l int) binSearchList {
	return binSearchList{size: l}
}

func (b binSearchList) length() int {
	var bin int = 1
	var start int = 0
	for base(b).shorterThan(start) < 0 {
		start = start + bin
		bin *= 2
	}
	return b.binSearch(start-bin/2, bin/2)
}

// search length from start to start+bin-1
func (b binSearchList) binSearch(start, bin int) int {
	//fmt.Printf("binSearch(%v, %v)\n", start, bin)
	if bin == 1 {
		return start
	}
	bin = bin / 2
	res := base(b).shorterThan(start + bin)
	switch {
	case res < 0:
		return b.binSearch(start+bin, bin)
	case res == 0:
		return start + bin
	default:
		return b.binSearch(start, bin)
	}
}

func main() {
}
