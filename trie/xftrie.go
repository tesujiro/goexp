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
	//debugf("Add(%v)\n", x)

	// 1 - search for x until following out oft trie
	for i := uint(0); i < bt.w; i++ {
		c := uint(x) >> (bt.w - i - 1) & 1

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
			val := number(uint(x) >> (bt.w - i))
			bt.t[i][val] = u.child[c]
		}
		u = u.child[c]
	}
	if exist {
		return false
	}
	u.x = x
	bt.t[bt.w][x] = u

	// 3 = add u to linked list
	u.child[0] = pred
	u.child[1] = pred.child[1]
	u.child[0].child[1] = u
	u.child[1].child[0] = u
	//debugf("%v => %v => %v\n", u.child[0].x, u.x, u.child[1].x)

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
	l := uint(0)  //law
	h := bt.w + 1 //hight
	u := bt.root
	loop := 0
	for h-l > 1 {
		i := (l + h) / 2
		p := x >> (bt.w - i)
		v, ok := bt.t[i][p]
		if !ok {
			debugf("N ")
			h = i
		} else {
			debugf("Y ")
			u = v
			l = i
		}
		loop++
	}
	debugf("\n")
	//debugf("loop=%v \n", loop)
	// found x
	if l == bt.w {
		//debugf("found u.x=%v\n", int(u.x))
		debugf("loop=%v FOUND %v\n", loop, x)
		return u.x
	}

	if u.jump == nil {
		debugf("loop=%v not found %v u.jump==nil \n", loop, x)
		return 0
	}
	// search for next value
	c := x >> (bt.w - l - 1) & 1
	var pred *node
	if c == 1 {
		pred = u.jump
	} else {
		//debugf("u=%#v\n", u)
		//debugf("u.jump=%#v\n", u.jump)
		pred = u.jump.child[0]
	}
	//if pred.right() == bt.dummy || pred.right() == nil {
	if pred.child[1] == bt.dummy {
		debugf("loop=%v, not found %v, return 0\n", loop, x)
		return 0
	} else {
		debugf("loop=%v, not found %v, return next %v, next?:%v\n", loop, x, pred.child[1].x, x < pred.child[1].x)
		return pred.child[1].x
	}
}
