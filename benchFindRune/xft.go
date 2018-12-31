package main

import "fmt"

/* X-Fast Trie */
const bitlen = 32

type node struct {
	x      interval
	parent *node
	child  [2]*node
	jump   *node
}

type xFastTrie struct {
	root  *node
	dummy *node
	w     uint             // bit length
	t     []map[rune]*node // hash
}

type XFT []*xFastTrie

func newXFT(ts []table) *XFT {
	ret := make(XFT, len(ts))
	for i, t := range ts {
		ret[i] = newXFastTrie(t)
	}
	return &ret
}

func (xft *XFT) find(r rune) bool {
	for _, t := range *xft {
		if t.find(r) {
			return true
		}
	}
	return false
}

func newXFastTrie(tbl table) *xFastTrie {
	dummy := &node{}
	dummy.child = [2]*node{dummy, dummy}
	t := make([]map[rune]*node, bitlen+1)
	for i := 0; i < len(t); i++ {
		t[i] = make(map[rune]*node)
		//t[i] = make(map[interval]*node, 1000)
	}

	xft := &xFastTrie{
		root:  &node{jump: dummy},
		dummy: dummy,
		w:     bitlen,
		t:     t,
	}

	for _, v := range tbl {
		xft.add(v)
	}
	//xft.print()

	return xft
}

func (xft *xFastTrie) getAll() (result []interval) {
	w := xft.dummy.child[1]
	for w != nil && w != xft.dummy {
		result = append(result, w.x)
		w = w.child[1]
	}
	return
}

func (xft *xFastTrie) print() {
	//fmt.Printf("xFastTrie=%v\n", xft.getAll())
	for i, x := range xft.getAll() {
		_ = i
		fmt.Printf("%x-%x ", x.first, x.last)
		if i%16 == 0 {
			fmt.Printf("\n")
		}
	}
}

func (xft *xFastTrie) add(x interval) bool {
	u := xft.root
	var pred *node // predecessor : jump連鎖上で追加すべきノードの一つ前のノード  see 3.
	exist := true

	// 1 - search for x until following out oft trie
	for i := uint(0); i < xft.w; i++ {
		c := uint(x.last) >> (xft.w - i - 1) & 1

		// if not found set pred
		if exist && u.child[c] == nil {
			exist = false
			if c == 0 { //left
				pred = u.jump.child[0]
			} else { //right
				pred = u.jump
			}

			u.jump = nil
		}

		// 2 - if not found add path to x
		if !exist {
			u.child[c] = &node{}
			u.child[c].parent = u
			val := rune(uint(x.last) >> (xft.w - i - 1))
			xft.t[i+1][val] = u.child[c]
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

	// 4 - walk back up, updating jump pointers
	for v := u.parent; v != nil; v = v.parent {
		if (v.child[0] == nil && (v.jump == nil || v.jump.x.last > x.last)) ||
			(v.child[1] == nil && (v.jump == nil || v.jump.x.last < x.last)) {
			v.jump = u
		} else {
			if v.jump == nil {
			}
		}
	}

	return true
}

//func (xft *xFastTrie) Find(x interval) interval {
func (xft *xFastTrie) find(r rune) bool {
	u := xft.root
	bot := uint(0)   // bottom
	top := xft.w + 1 // top
	for top-bot > 1 {
		mid := (bot + top) >> 1
		//p := x >> (xft.w - mid)
		//if v, ok := xft.t[mid][p]; !ok {
		if v, ok := xft.t[mid][r>>(xft.w-mid)]; !ok {
			top = mid
		} else {
			u = v
			bot = mid
		}
	}
	// found x
	if bot == xft.w {
		fmt.Println("Found")
		return true
		//return u.x
	}

	// search for next value
	c := r >> (xft.w - bot - 1) & 1
	var pred *node
	if c == 1 {
		pred = u.jump
	} else {
		pred = u.jump.child[0]
	}
	if pred.child[1] == xft.dummy {
		return false
	} else {
		//fmt.Printf("r=%x\tpred.x.first=%x\tpred.x.last=%x\n", r, pred.child[1].x.first, pred.child[1].x.last)
		//fmt.Println("return false")
		return pred.child[1].x.first <= r
	}
}
