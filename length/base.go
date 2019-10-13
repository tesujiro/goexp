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

func (b base) shorterThan(length int) int {
	if b.size < length {
		return 1
	} else if b.size > length {
		return -1
	}
	return 0
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
