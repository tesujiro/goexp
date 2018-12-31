package main

import "fmt"

func main() {
	tables := []table{private, nonprint, combining, doublewidth, ambiguous, emoji, notassigned, neutral}
	fmt.Println("\n== Binary Search ==")
	find(newBinary(tables))
	//find(newHashtable(tables))
	fmt.Println("\n== X-Fast Trie ==")
	find(newXFT(tables))
}

func find(f finder) {
	runes := []rune{'a', 'あ', '🍺', '你', 'ｱ'}

	for _, r := range runes {
		fmt.Printf("%c inTable=%v\n", r, f.find(r))
	}
}

type finder interface {
	find(rune) bool
}

/* Hashtable search */
type hashtable struct {
	hash []map[interval]struct{}
}

func newHashtable(ts []table) hashtable {
	ret := make([]map[interval]struct{}, len(ts))
	for i, t := range ts {
		h := make(map[interval]struct{}, len(t))
		for _, v := range t {
			h[v] = struct{}{}
		}
		ret[i] = h
	}
	return hashtable{hash: ret}
}

// NG!! imposibble to implement
func (ht hashtable) find(r rune) bool {
	for _, m := range ht.hash {
		if _, ok := m[interval{first: r}]; ok { // NG!
			return true
		}
	}
	return false
}
