package main

import (
	"math"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func init() {
}

func uniq(t []number) (result []number) {
	result = []number{}
	var prev number
	for i, v := range t {
		if i == 0 || prev != v {
			result = append(result, v)
		}
		prev = v
	}
	return
}

func testBinaryTrie_Add(t *testing.T) {
	// prepare test data
	max := int(math.Pow(2, bitlen))
	var table []number
	const numbers = 51200

	table = []number{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numbers; i++ {
		table = append(table, number(rand.Intn(max)))
	}

	//
	bt := newBinaryTrie()
	for _, v := range table {
		b := bt.Add(v)
		_ = b
	}
	//bt.Print()
	result := bt.GetAll()

	sort.Slice(table, func(i, j int) bool { return table[i] < table[j] })
	table = uniq(table)
	if !reflect.DeepEqual(table, result) {
		t.Fatalf("failed add %#v", result)
	}

}

func TestBinaryTrie_Add(t *testing.T) {
	// prepare test data
	for i := 0; i < 3; i++ {
		testBinaryTrie_Add(t)
	}
}

func BenchmarkBinaryTrie_Add(b *testing.B) {
	max := int(math.Pow(2, bitlen))
	table := []number{}
	//rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		table = append(table, number(rand.Intn(max)))
	}

	bt := newBinaryTrie()
	b.ResetTimer()
	for _, i := range table {
		bt.Add(number(i))
	}
}
