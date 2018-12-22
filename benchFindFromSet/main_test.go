package main

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func testFind(f finder, t *testing.T) {
	list := numbers{number(100), number(1), number(5), number(30), number(-1), number(-10)}
	test_table := numbers{number(1), number(0), number(7), number(101), number(-10)}
	expected := []bool{true, false, false, false, true}

	f.set(list)
	for i, n := range test_table {
		actual := f.find(n)
		if actual != expected[i] {
			t.Errorf("n=%v expected=%v actual=%v\n", n, expected[i], actual)
		}
	}
}

func TestFind_Binary(t *testing.T) {
	testFind(newBinary(), t)
}

func TestFind_Hashtable(t *testing.T) {
	testFind(newHashtable(), t)
}

func benchFind(f finder, elements int, b *testing.B) {
	//rand.Seed(time.Now().UnixNano())
	max := int(math.Pow(2, width))
	list := make(numbers, elements)
	test_table := make(numbers, b.N)

	for i := 0; i < elements; i++ {
		list[i] = number(rand.Intn(max))
	}

	for i := 0; i < b.N; i++ {
		test_table[i] = number(rand.Intn(max))
	}

	f.set(list)
	b.ResetTimer()
	for _, n := range test_table {
		f.find(n)
	}
	b.StopTimer()
}

func benchFindLoop(f func() finder, b *testing.B) {
	es := make([]int, 20)
	for i := 0; i < len(es); i++ {
		es[i] = 2 << uint(i)
	}
	for i, e := range es {
		b.Run(fmt.Sprintf("FindFrom2^%v", i), func(b *testing.B) {
			//benchFind(newBinary(), e, b)
			benchFind(f(), e, b)
		})
	}
}

func BenchmarkBinary(b *testing.B) {
	benchFindLoop(func() finder {
		return finder(newBinary())
	}, b)
}

func BenchmarkHashtable(b *testing.B) {
	benchFindLoop(func() finder {
		return finder(newHashtable())
	}, b)
}
