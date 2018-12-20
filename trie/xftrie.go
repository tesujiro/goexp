package main

//var Empty struct{}

type xFastTrie struct {
	root  *node
	dummy *node
	w     uint               // bit length
	t     []map[number]*node // hash
}

func newXFastTrie() *xFastTrie {
	dummy := &node{}
	dummy.child = [2]*node{dummy, dummy}
	t := make([]map[number]*node, bitlen+1)
	for i := 0; i < len(t); i++ {
		t[i] = make(map[number]*node)
		//t[i] = make(map[number]*node, 1000)
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
	debugf("xFastTrie=%v\n", bt.GetAll())
}

func (bt *xFastTrie) Add(x number) bool {
	u := bt.root
	var pred *node // predecessor : jump連鎖上で追加すべきノードの一つ前のノード  see 3.
	exist := true

	// 1 - search for x until following out oft trie
	for i := uint(0); i < bt.w; i++ {
		c := uint(x) >> (bt.w - i - 1) & 1

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
			val := number(uint(x) >> (bt.w - i - 1))
			bt.t[i+1][val] = u.child[c]
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
		if (v.child[0] == nil && (v.jump == nil || v.jump.x > x)) ||
			(v.child[1] == nil && (v.jump == nil || v.jump.x < x)) {
			v.jump = u
		} else {
			if v.jump == nil {
			}
		}
	}

	return true
}

func (bt *xFastTrie) Find(x number) number {
	u := bt.root
	bot := uint(0)  // bottom
	top := bt.w + 1 // top
	for top-bot > 1 {
		mid := (bot + top) >> 1
		p := x >> (bt.w - mid)
		if v, ok := bt.t[mid][p]; !ok {
			top = mid
		} else {
			u = v
			bot = mid
		}
	}
	// found x
	if bot == bt.w {
		return u.x
	}

	// search for next value
	c := x >> (bt.w - bot - 1) & 1
	var pred *node
	if c == 1 {
		pred = u.jump
	} else {
		pred = u.jump.child[0]
	}
	if pred.child[1] == bt.dummy {
		return 0
	} else {
		return pred.child[1].x
	}
}
