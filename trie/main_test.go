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

func contains(t []number, i number) bool {
	for _, v := range t {
		if v == i {
			return true
		}
	}
	return false
}

func testBinaryTrie(bt trie, t *testing.T) {
	// prepare test data
	max := int(math.Pow(2, bitlen))

	rand.Seed(time.Now().UnixNano())
	table := []number{}
	for i := 0; i < add_count; i++ {
		table = append(table, number(rand.Intn(max)))
	}
	test_table := []number{}
	for i := 0; i < find_count; i++ {
		test_table = append(test_table, number(rand.Intn(max)))
	}

	//

	// Add
	for _, v := range table {
		b := bt.Add(v)
		_ = b
	}
	//bt.Print()
	result := bt.GetAll()
	debugf("trie=%v\n", result)

	sort.Slice(table, func(i, j int) bool { return table[i] < table[j] })
	table = uniq(table)
	if !reflect.DeepEqual(table, result) {
		t.Fatalf("failed add %#v", result)
	}

	// Find
	for _, v := range test_table {
		n := bt.Find(v)
		b := contains(table, v)
		if n != v && b || n == v && !b {
			t.Fatalf("failed find %#v contains=%v", int(v), b)
		}
		/*
			if b {
				t.Logf("found %#v", int(v))
			}
		*/
	}
}

func TestBinaryTrie(t *testing.T) {
	// prepare test data
	for i := 0; i < 1; i++ {
		testBinaryTrie(newBinaryTrie(), t)
	}
}

func TestXFastTrie(t *testing.T) {
	// prepare test data
	for i := 0; i < 1; i++ {
		testBinaryTrie(newXFastTrie(), t)
	}
}

func benchmarkTrie_Add(bt trie, b *testing.B) {
	max := int(math.Pow(2, bitlen))
	table := []number{}
	//rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		table = append(table, number(rand.Intn(max)))
	}

	b.ResetTimer()
	for _, i := range table {
		bt.Add(number(i))
	}
}

func BenchmarkBinaryTrie_Add(b *testing.B) {
	benchmarkTrie_Add(newBinaryTrie(), b)
}

func BenchmarkXFastTrie_Add(b *testing.B) {
	benchmarkTrie_Add(newXFastTrie(), b)
}

func benchmarkTrie_Find(bt trie, b *testing.B) {
	max := int(math.Pow(2, bitlen))
	table := []number{}
	test_table := []number{}
	//rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		table = append(table, number(rand.Intn(max)))
		test_table = append(test_table, number(rand.Intn(max)))
	}

	for _, i := range table {
		bt.Add(number(i))
	}

	b.ResetTimer()
	for _, i := range test_table {
		bt.Find(number(i))
	}
	//b.StopTimer()
}

func BenchmarkBinaryTrie_Find(b *testing.B) {
	benchmarkTrie_Find(newBinaryTrie(), b)
}

func BenchmarkXFastTrie_Find(b *testing.B) {
	benchmarkTrie_Find(newXFastTrie(), b)
}
