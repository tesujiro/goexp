package main

import "fmt"

type number uint32

const bitlen = 32

const debug = false

// for test
const add_count = 512
const find_count = 5120

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

func debugf(format string, a ...interface{}) (n int, err error) {
	if debug {
		return fmt.Printf(format, a...)
	}
	return 0, nil
}
