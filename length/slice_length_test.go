package main

import "testing"

func TestMain(t *testing.T) {
	expected := 4120
	s := make(naiveSlice, expected)
	actual := s.length()
	if actual != expected {
		t.Errorf("want: %v actual: %v", expected, actual)
	}
	b := make(binarySlice, expected)
	actual = b.length()
	if actual != expected {
		t.Errorf("want: %v actual: %v", expected, actual)
	}

}

func BenchmarkLength(b *testing.B) {
	var l int
	b.Run("naiveSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := make(naiveSlice, i)
			l = s.length()
		}
	})
	b.Run("binSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := make(binarySlice, i)
			l = s.length()
		}
	})
	_ = l
}
