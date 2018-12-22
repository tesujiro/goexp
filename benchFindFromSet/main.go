package main

import (
	"fmt"
	"sort"
)

type number int32

const width = 32

type numbers []number

func (ns numbers) Len() int {
	return len(ns)
}

func (ns numbers) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}

func (ns numbers) Less(i, j int) bool {
	return ns[i] < ns[j]
}

type finder interface {
	set(numbers)
	find(number) bool
}

type binary struct {
	orderedList numbers
}

func newBinary() *binary {
	return &binary{}
}

func (b *binary) set(s numbers) {
	sort.Sort(numbers(s))
	ss := make(numbers, len(s))
	copy(ss, s)
	b.orderedList = ss
}

func (b *binary) find(n number) bool {
	//fmt.Printf("s=%v\n", b.orderedList)
	bot := 0
	top := len(b.orderedList) - 1
	for bot <= top {
		mid := (bot + top) >> 1
		//fmt.Printf("b.orderedList[mid]=%v n=%v\n", b.orderedList[mid], n)
		switch {
		case b.orderedList[mid] < n:
			bot = mid + 1
		case b.orderedList[mid] > n:
			top = mid - 1
		default:
			return true
		}
	}
	return false
}

type hashtable struct {
	hash map[number]struct{}
}

func newHashtable() *hashtable {
	return &hashtable{}
}
func (ht *hashtable) set(ns numbers) {
	h := make(map[number]struct{}, len(ns))
	for _, v := range ns {
		h[v] = struct{}{}
	}
	ht.hash = h
}

func (ht *hashtable) find(n number) bool {
	if _, ok := ht.hash[n]; ok {
		return true
	}
	return false
}

func find(f finder) {
	list := numbers{number(100), number(1), number(5), number(30), number(-1), number(-10)}
	test_table := numbers{number(1), number(0), number(7), number(101), number(-10)}

	f.set(list)
	for _, n := range test_table {
		fmt.Printf("%v inTable=%v\n", n, f.find(n))
	}
}

func main() {
	find(newBinary())
	find(newHashtable())
}
