package main

import (
	"math/rand"
	"testing"
)

var tables = []table{private, nonprint, combining, doublewidth, ambiguous, emoji, notassigned, neutral}

func TestBinarySearch(t *testing.T) {
	testFind(newBinary(tables), t)
}

func TestXFastTrie(t *testing.T) {
	testFind(newXFT(tables), t)
}

func testFind(f finder, t *testing.T) {
	runes := []rune{'a', 'あ', '🍺', '你', 'ｱ'}
	expected := []bool{false, true, true, true, false}

	for i, r := range runes {
		//fmt.Printf("%c inTable=%v\n", r, f.find(r))
		actual := f.find(r)
		if actual != expected[i] {
			t.Errorf("error r=%c\n actual=%v expected=%v\n", r, actual, expected[i])
		}
	}
}

func BenchmarkBinary_Find(b *testing.B) {
	benchmark_Find(func() finder {
		return newBinary(tables)
	}, b)
}

func BenchmarkXFT_Find(b *testing.B) {
	benchmark_Find(func() finder {
		return newXFT(tables)
	}, b)
}

func benchmark_Find(f func() finder, b *testing.B) {
	test_table := make([]rune, b.N)
	for i := 0; i < len(test_table); i++ {
		test_table[i] = rune(rand.Intn(256 * 256 * 256 * 256))
	}
	finder := f()
	b.ResetTimer()
	for _, v := range test_table {
		finder.find(v)
	}
}
