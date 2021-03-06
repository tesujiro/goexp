package main

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

type baseWithLimitedSizeInLimitedFailure struct {
	baseWithLimitedSize
	fails int
}

func (b base) shorterThan(length int) int {
	if b.size < length {
		return 1
	} else if b.size > length {
		return -1
	}
	return 0
}

func (b base) bruteForce(start, end int) int {
	var i int
	if end < 0 {
		for i = start; b.shorterThan(i) < 0; i++ {
		}
	} else {
		for i = start; b.shorterThan(i) < 0 && i < end; i++ {
		}
	}
	return i
}

func (b base) binSearch(start int, bit uint) int {
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
