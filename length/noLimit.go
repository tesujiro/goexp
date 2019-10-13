package main

type naiveList base

func newNaiveList(l int) naiveList {
	return naiveList{size: l}
}

func (n naiveList) length() int {
	return base(n).bruteForce(0, -1)
}

type binSearchList base

func newBinSearchList(l int) binSearchList {
	return binSearchList{size: l}
}

func (b binSearchList) length() int {
	var bit uint = 0
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
