package main

import (
	"math/rand"
	"testing"
)

func TestMain(t *testing.T) {
	test := func(expected int) {
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
	}
	test(0)
	for i := 0; i < 100; i++ {
		test(rand.Intn(1000000))
	}

}

func BenchmarkLength(b *testing.B) {
	var l int
	var s lengther
	b.Run("naiveList", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s = newNaiveList(i)
			l = s.length()
		}
	})
	b.Run("binSearchList", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s = newBinSearchList(i)
			l = s.length()
		}
	})
	_ = l
}
