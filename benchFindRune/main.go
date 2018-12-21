package main

import "fmt"

type finder interface {
	find(rune) bool
}

type binary struct {
	tables []table
}

func newBinary(ts []table) binary {
	return binary{tables: ts}
}

func (b binary) find(r rune) bool {
	for _, t := range b.tables {
		if b.findTable(r, t) {
			return true
		}
	}
	return false
}

func (b binary) findTable(r rune, t table) bool {
	// func (t table) IncludesRune(r rune) bool {
	if r < t[0].first {
		return false
	}

	bot := 0
	top := len(t) - 1
	for top >= bot {
		mid := (bot + top) / 2

		switch {
		case t[mid].last < r:
			bot = mid + 1
		case t[mid].first > r:
			top = mid - 1
		default:
			return true
		}
	}

	return false
}

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

func find(f finder) {
	runes := []rune{'a', 'あ', '🍺', '你', 'ｱ'}

	for _, r := range runes {
		fmt.Printf("%c inTable=%v\n", r, f.find(r))
	}
}

func main() {
	tables := []table{private, nonprint, combining, doublewidth, ambiguous, emoji, notassigned, neutral}
	find(newBinary(tables))
	find(newHashtable(tables))
}
