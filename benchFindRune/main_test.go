package main

import (
	"testing"
)

var tables = []table{private, nonprint, combining, doublewidth, ambiguous, emoji, notassigned, neutral}

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

func TestBinarySearch(t *testing.T) {
	testFind(newBinary(tables), t)
}

func TestXFastTrie(t *testing.T) {
	testFind(newXFT(tables), t)
}
