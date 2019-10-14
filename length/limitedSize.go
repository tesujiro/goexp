package main

import (
	"math"
)

type binSearchListWithLimitedSize baseWithLimitedSize

func newBinSearchListWithLimitedSize(l, m int) binSearchListWithLimitedSize {
	return binSearchListWithLimitedSize{base{l}, m}
}

func (b binSearchListWithLimitedSize) length() int {
	n := uint(math.Ceil(math.Log2(float64(b.max))))
	return baseWithLimitedSize(b).binSearch(0, n)
}

type searchListInLtdFails baseWithLimitedSizeInLimitedFailure

func newSearchListWithLimitedFails(l, m, f int) searchListInLtdFails {
	return searchListInLtdFails{baseWithLimitedSize{base{l}, m}, f}
}

func (l searchListInLtdFails) length() int {
	/*
		start := 0
		bit := uint(math.Ceil(math.Log2(float64(l.baseWithLimitedSize.max))))
		for f := l.fails; f > 1; {
			bit--
			if bit == 0 {
				return start
			}
			next := start + 1<<bit
			res := l.baseWithLimitedSize.shorterThan(next)
			switch {
			case res < 0:
				start = next
				f--
			case res == 0:
				return next
			}
		}
	*/
	next := func(max int) int {
		// n*(n+1)/2 = max
		// n*n + n - 2*max = 0
		// n = -1/2 + Sqrt(1+4*2*max)/2
		n := int(math.Ceil(-0.5 + math.Sqrt(float64(1+8*max))/2))
		return n
	}

	start := 0
	end := l.baseWithLimitedSize.max
	for f := l.fails; f > 0; {
		next := start + next(end-start)
		res := l.baseWithLimitedSize.shorterThan(next)
		switch {
		case res < 0:
			start = next
			//fmt.Printf("OK rest=%v\tstart=%v\tend=%v\n", f, start, end)
		case res == 0:
			return next
		case res > 0:
			end = next
			f--
			//fmt.Printf("NG rest=%v\tstart=%v\tend=%v\n", f, start, end)
		}
	}
	//fmt.Printf("start=%v\tend=%v\n", start, end)
	return l.base.bruteForce(start, end)
}
