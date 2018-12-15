package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type number uint8

const bitlen = 8
const numbers = 16

var table []number

func init() {
	max := int(math.Pow(2, bitlen))

	table = []number{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numbers; i++ {
		table = append(table, number(rand.Intn(max)))
	}
}

type node struct {
	x      number
	parent *node
	left   *node
	right  *node
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
	dummy.left = dummy
	dummy.right = dummy

	return &binaryTrie{
		root:  &node{jump: dummy},
		dummy: dummy,
		w:     bitlen,
	}
}

func (bt *binaryTrie) GetAll() (result []number) {
	w := bt.dummy.right
	for w != nil && w != bt.dummy {
		result = append(result, w.x)
		w = w.right
	}
	return
}

func (bt *binaryTrie) Print() {
	fmt.Printf("binaryTrie=%v\n", bt.GetAll())
}

func (bt *binaryTrie) Add(x number) bool {
	u := bt.root
	var i, c uint

	// 1 - search for x until following out oft trie
	for i = 0; i < bt.w; i++ {
		c = uint(x) >> (bt.w - i - 1) & 1
		if u == nil {
			break
		}
		if c == 0 {
			if u.left == nil {
				break
			}
			u = u.left
		} else {
			if u.right == nil {
				break
			}
			u = u.right
		}
	}
	if i == bt.w {
		// bt already has x
		return false
	}
	var pred *node
	if c == 1 { //right
		pred = u.jump
	} else { //left
		pred = u.jump.left
	}
	/*
		if pred == nil {
			pred = bt.dummy
		}
	*/
	u.jump = nil

	// 2 - add path to x
	for ; i < bt.w; i++ {
		c = uint(x) >> (bt.w - i - 1) & 1
		if c == 0 {
			u.left = &node{}
			u.left.parent = u
			u = u.left
		} else {
			u.right = &node{}
			u.right.parent = u
			u = u.right
		}
		/*
			var n **node
			if c == 0 {
				n = &u.left
			} else {
				n = &u.right
			}
			*n = &node{}
			(*n).parent = u
			u = *n
		*/
	}
	u.x = x

	// 3 = add u to linked list
	u.left = pred
	u.right = pred.right
	u.left.right = u
	u.right.left = u
	//fmt.Printf("%v => %v => %v\n", u.left.x, u.x, u.right.x)

	// 4 - walk back up, updating jump pointers
	v := u.parent
	for v != nil {
		if (v.left == nil && (v.jump == nil || v.jump.x > x)) ||
			(v.right == nil && (v.jump == nil || v.jump.x < x)) {
			v.jump = u
		}
		v = v.parent
	}

	return true
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
