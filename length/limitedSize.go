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
