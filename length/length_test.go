package main

import (
	"math/rand"
	"testing"
)

func TestMain(t *testing.T) {
	test := func(expected, max int) {
		n := newNaiveList(expected)
		actual := n.length()
		if actual != expected {
			t.Errorf("want: %v actual: %v", expected, actual)
		}
		bs := newBinSearchList(expected)
		actual = bs.length()
		if actual != expected {
			t.Errorf("want: %v actual: %v", expected, actual)
		}
		bs2 := newBinSearchListWithLimitedSize(expected, max)
		actual = bs2.length()
		if actual != expected {
			t.Errorf("want: %v actual: %v", expected, actual)
		}
		lf := newSearchListWithLimitedFails(expected, max, 3)
		actual = lf.length()
		if actual != expected {
			t.Errorf("want: %v actual: %v", expected, actual)
		}
	}
	//test(6142, 10000)
	//test(9432, 10000)
	test(0, 0)
	for i := 0; i < 10; i++ {
		max := rand.Intn(1000000)
		num := rand.Intn(max)
		test(num, max)
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
			s = newBinSearchListWithLimitedSize(rand.Intn(i+1), i) // TODO; not the same test
			l = s.length()
		}
	})
	b.Run("binSearchListInLimitedFails", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s = newSearchListWithLimitedFails(rand.Intn(i+1), i, rand.Intn(10)) // TODO; not the same test
			l = s.length()
		}
	})
	_ = l
}
