package main

import "fmt"

type number uint32

const bitlen = 2
const add_count = 1
const find_count = 6

type node struct {
	x      number
	parent *node
	child  [2]*node
	jump   *node
}

type trie interface {
	GetAll() []number
	Print()
	Add(number) bool
	Find(number) number
}

const debug = true

func debugf(format string, a ...interface{}) (n int, err error) {
	if debug {
		return fmt.Printf(format, a...)
	}
	return 0, nil
}
