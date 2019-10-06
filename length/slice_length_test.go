package main

import "testing"

func TestMain(t *testing.T) {
	expected := 4120
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

func BenchmarkLength(b *testing.B) {
	var l int
	b.Run("naiveList", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := newNaiveList(i)
			l = s.length()
		}
	})
	b.Run("binSearchList", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := newBinSearchList(i)
			l = s.length()
		}
	})
	_ = l
}
