package main

import (
	"math/rand"
	"testing"
)

func TestMain(t *testing.T) {
	test := func(expected, max, fails int) {
		n := newNaiveList(expected)
		actual := n.length()
		if actual != expected {
			t.Errorf("error naiveList want: %v actual: %v", expected, actual)
		}
		bs := newBinSearchList(expected)
		actual = bs.length()
		if actual != expected {
			t.Errorf("error binSearchList want: %v actual: %v", expected, actual)
		}
		bs2 := newBinSearchListWithLimitedSize(expected, max)
		actual = bs2.length()
		if actual != expected {
			t.Errorf("error binSearchListWithLimitedSize want: %v actual: %v max:%v", expected, actual, max)
		}
		lf := newSearchListWithLimitedFails(expected, max, fails)
		actual = lf.length()
		if actual != expected {
			t.Errorf("error searchListWithLimitedFials want: %v actual: %v max:%v fails:%v", expected, actual, max, fails)
		}
	}
	test(6142, 10000, 1)
	test(9432, 10000, 9)
	test(41593, 100000, 9)
	test(0, 1, 0)
	for i := 0; i < 1000; i++ {
		max := rand.Intn(10000000) + 1
		num := rand.Intn(max)
		fails := rand.Intn(10)
		//fmt.Printf("test(num=%v,max=%v)\n", num, max)
		test(num, max, fails)
	}
}

func BenchmarkLength(b *testing.B) {
	var l int
	var s lengther
	b.Run("naiveListWithoutMaxSize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s = newNaiveList(i)
			l = s.length()
		}
	})
	b.Run("binSearchListWithoutMaxSize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s = newBinSearchList(i)
			l = s.length()
		}
	})
	b.Run("binSearchListWithLimitedSize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s = newBinSearchListWithLimitedSize(rand.Intn(i+1), i+1) // TODO; not the same test
			l = s.length()
		}
	})
	b.Run("binSearchListInLimitedFails", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s = newSearchListWithLimitedFails(rand.Intn(i+1), i+1, rand.Intn(10)) // TODO; not the same test
			l = s.length()
		}
	})
	_ = l
}
