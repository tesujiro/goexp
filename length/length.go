package main

import (
	"math"
	"math/rand"

	"github.com/pkg/profile"
)

type lengther interface {
	length() int
}

type base struct {
	size int
}

type base2 struct {
	base
	max int
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

type binSearchList2 base2

func newBinSearchList2(l, m int) binSearchList2 {
	return binSearchList2{base{l}, m}
}

func (b binSearchList2) length() int {
	n := int(math.Ceil(math.Log2(float64(b.max))))
	return b.binSearch(0, n)
}

func (b binSearchList2) binSearch(start, bit int) int {
	if bit == 0 {
		return start
	}
	bit--
	res := (b.base).shorterThan(start + 1<<bit)
	switch {
	case res < 0:
		return b.binSearch(start+1<<bit, bit)
	case res == 0:
		return start + 1<<bit
	default:
		return b.binSearch(start, bit)
	}
}

func main() {
	var l int
	N := 10000000
	defer profile.Start(profile.ProfilePath(".")).Stop()
	for i := 0; i < N; i++ {
		rnd := rand.Intn(N)
		s := newBinSearchList2(rnd, rnd+1)
		l = s.length()
	}
	_ = l
}
