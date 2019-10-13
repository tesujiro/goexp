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

type baseWithLimitedSize struct {
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

func (b base) binSearch(start, bit int) int {
	if bit == 0 {
		return start
	}
	bit--
	next := start + 1<<bit
	res := b.shorterThan(next)
	switch {
	case res < 0:
		return b.binSearch(next, bit)
	case res == 0:
		return next
	default:
		return b.binSearch(start, bit)
	}
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
	var bit int = 0
	var start int = 0
	for base(b).shorterThan(start) < 0 {
		bit++
		start += 1 << bit
	}
	if bit == 0 {
		return 0
	}
	return base(b).binSearch(start-1<<bit, bit)
}

type binSearchListWithLimitedSize baseWithLimitedSize

func newBinSearchListWithLimitedSize(l, m int) binSearchListWithLimitedSize {
	return binSearchListWithLimitedSize{base{l}, m}
}

func (b binSearchListWithLimitedSize) length() int {
	n := int(math.Ceil(math.Log2(float64(b.max))))
	return baseWithLimitedSize(b).binSearch(0, n)
}

func main() {
	var l int
	N := 10000000
	defer profile.Start(profile.ProfilePath(".")).Stop()
	for i := 0; i < N; i++ {
		rnd := rand.Intn(N)
		s := newBinSearchListWithLimitedSize(rnd, rnd+1)
		l = s.length()
	}
	_ = l
}
