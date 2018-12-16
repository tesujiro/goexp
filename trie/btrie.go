package main

import (
	"fmt"
)

type number uint8

const bitlen = 8

type node struct {
	x      number
	parent *node
	child  [2]*node
	jump   *node
}

//func newNode() *node {
//return &node{}
//}

type binaryTrie struct {
	root  *node
	dummy *node
	w     uint // bit length
}

func newBinaryTrie() *binaryTrie {
	dummy := &node{}
	dummy.child = [2]*node{dummy, dummy}

	return &binaryTrie{
		root:  &node{jump: dummy},
		dummy: dummy,
		w:     bitlen,
	}
}

func (bt *binaryTrie) GetAll() (result []number) {
	w := bt.dummy.child[1]
	for w != nil && w != bt.dummy {
		result = append(result, w.x)
		w = w.child[1]
	}
	return
}

func (bt *binaryTrie) Print() {
	fmt.Printf("binaryTrie=%v\n", bt.GetAll())
}

func (bt *binaryTrie) Add(x number) bool {
	u := bt.root
	var i, c uint
	var pred *node
	//already := false
	//fmt.Printf("Add(%v)\n", x)

	// 1 - search for x until following out oft trie
	for ; ; i++ {
		if i == bt.w {
			// bt already has x
			return false
		}
		c = uint(x) >> (bt.w - i - 1) & 1
		if u == nil || u.child[c] == nil {
			//already = true
			if c == 0 { //right
				pred = u.jump.child[c]
			} else { //left
				pred = u.jump
			}
			u.jump = nil
			break
		}
		u = u.child[c]
	}

	// 2 - add path to x
	//for ; i < bt.w; i++ {
	for ; ; i++ {
		if i == bt.w {
			break
		}
		c = uint(x) >> (bt.w - i - 1) & 1
		u.child[c] = &node{}
		u.child[c].parent = u
		u = u.child[c]
	}
	u.x = x

	// 3 = add u to linked list
	u.child[0] = pred
	u.child[1] = pred.child[1]
	u.child[0].child[1] = u
	u.child[1].child[0] = u
	//fmt.Printf("%v => %v => %v\n", u.child[0].x, u.x, u.child[1].x)

	// 4 - walk back up, updating jump pointers
	v := u.parent
	for v != nil {
		if (v.child[0] == nil && (v.jump == nil || v.jump.x > x)) ||
			(v.child[1] == nil && (v.jump == nil || v.jump.x < x)) {
			v.jump = u
		}
		v = v.parent
	}

	return true
}

/*
func (e1 number) Compare(e2 ch01.Comparable) int {
	return int(e1 - e2.(number))
}

func (e1 number) HashCode() uint {
	return uint(e1)
}

func testCh13() {
	s := ch13.NewBinaryTrie()
	for _, v := range table {
		b := s.Add(v)
		_ = b
		//fmt.Printf("Add(%v)=>%v\n", v, b)
	}
	s.Print()
}

func main() {
	bt := newBinaryTrie()
	for _, v := range table {
		b := bt.Add(v)
		_ = b
		//fmt.Printf("Add(%v)=>%v\n", v, b)
	}
	bt.Print()

}
*/
