package main

import (
	"fmt"
)

/*
type number uint16

const bitlen = 16

type node struct {
	x      number
	parent *node
	child  [2]*node
	jump   *node
}
*/

var Empty struct{}

type xFastTrie struct {
	root  *node
	dummy *node
	w     uint                  // bit length
	t     []map[number]struct{} // hash
}

func newXFastTrie() *xFastTrie {
	dummy := &node{}
	dummy.child = [2]*node{dummy, dummy}
	t := make([]map[number]struct{}, bitlen+1)
	for i := 0; i < len(t); i++ {
		t[i] = make(map[number]struct{})
	}

	return &xFastTrie{
		root:  &node{jump: dummy},
		dummy: dummy,
		w:     bitlen,
		t:     t,
	}
}

func (bt *xFastTrie) GetAll() (result []number) {
	w := bt.dummy.child[1]
	for w != nil && w != bt.dummy {
		result = append(result, w.x)
		w = w.child[1]
	}
	return
}

func (bt *xFastTrie) Print() {
	fmt.Printf("xFastTrie=%v\n", bt.GetAll())
}

func (bt *xFastTrie) Add(x number) bool {
	u := bt.root
	var pred *node // predecessor : jump連鎖上で追加すべきノードの一つ前のノード  see 3.
	exist := true
	//fmt.Printf("Add(%v)\n", x)

	// 1 - search for x until following out oft trie
	bt.t[0][x] = Empty
	for i := uint(0); i < bt.w; i++ {
		c := uint(x) >> (bt.w - i - 1) & 1

		val := number(uint(x) >> (bt.w - i))
		bt.t[i+1][val] = Empty
		// if not found set pred
		if exist && u.child[c] == nil {
			exist = false
			if c == 0 { //right
				pred = u.jump.child[0]
			} else { //left
				pred = u.jump
			}
			u.jump = nil
		}

		// 2 - if not found add path to x
		if !exist {
			u.child[c] = &node{}
			u.child[c].parent = u
		}
		u = u.child[c]
	}
	if exist {
		return false
	}
	u.x = x

	// 3 = add u to linked list
	u.child[0] = pred
	u.child[1] = pred.child[1]
	u.child[0].child[1] = u
	u.child[1].child[0] = u
	//fmt.Printf("%v => %v => %v\n", u.child[0].x, u.x, u.child[1].x)

	// 4 - walk back up, updating jump pointers
	for v := u.parent; v != nil; v = v.parent {
		if (v.child[0] == nil && (v.jump == nil || v.jump.x > x)) ||
			(v.child[1] == nil && (v.jump == nil || v.jump.x < x)) {
			v.jump = u
		}
	}

	return true
}

func (bt *xFastTrie) Find(x number) number {
	u := bt.root
	for i := uint(0); i < bt.w; i++ {
		c := uint(x) >> (bt.w - i - 1) & 1
		if u.child[c] == nil {
			// not founc & search for next value
			if c == 0 {
				u = u.jump
			} else {
				u = u.jump.child[1]
			}
			if u == bt.dummy {
				return 0
			} else {
				return u.x
			}
		}
		u = u.child[c]
	}
	// found x
	return u.x
}
