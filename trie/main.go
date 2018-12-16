package main

type number uint16

const bitlen = 16

type node struct {
	x      number
	parent *node
	child  [2]*node
	jump   *node
}
