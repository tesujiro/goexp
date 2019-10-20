package main

import (
	"math"
)

type binSearchListWithLimitedSize baseWithLimitedSize

func newBinSearchListWithLimitedSize(l, m int) binSearchListWithLimitedSize {
	return binSearchListWithLimitedSize{base{l}, m}
}

func (b binSearchListWithLimitedSize) length() int {
	var n uint = 0
	for ; 1<<n < b.max; n++ {
	}

	return baseWithLimitedSize(b).binSearch(0, n)
}

type searchListInLtdFails baseWithLimitedSizeInLimitedFailure

func newSearchListWithLimitedFails(l, m, f int) searchListInLtdFails {
	return searchListInLtdFails{baseWithLimitedSize{base{l}, m}, f}
}

func (l searchListInLtdFails) length() int {
	// while fail > 1 do binary search.
	start := 0
	bit := uint(math.Ceil(math.Log2(float64(l.baseWithLimitedSize.max))))
	for f := l.fails; f > 1; {
		bit--
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

	summationToN := func(max int) int {
		// n*(n+1)/2 = max
		// n*n + n - 2*max = 0
		// n = -1/2 + Sqrt(1+4*2*max)/2
		n := int(math.Ceil(-0.5 + math.Sqrt(float64(1+8*max))/2))
		return n
	}
	next := summationToN(1 << bit)
	for {
		res := l.baseWithLimitedSize.shorterThan(start + next)
		switch {
		case res < 0:
			start = start + next
			next--
		case res == 0:
			return start + next
		case res > 0:
			return l.base.bruteForce(start, start+next)
		}
	}
}
