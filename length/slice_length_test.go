package main

import "testing"

func TestMain(t *testing.T) {
	expected := 4120
	s := make(naiveSlice, expected)
	actual := s.length()
	if actual != expected {
		t.Errorf("want: %v actual: %v", expected, actual)
	}
}

func BenchmarkLength(b *testing.B) {
	var l int
	for i := 0; i < b.N; i++ {
		s := make(naiveSlice, i)
		l = s.length()
	}
	_ = l
}
