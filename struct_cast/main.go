package main

import "fmt"

type parent struct {
	col interface{}
	i   int
	s   string
}

type childInt struct {
	parent
	col int
}

type childStruct struct {
	parent
	col string
}

func main() {
	p1 := parent{col: 1, i: 100, s: "abc"}
	fmt.Printf("p1=%#v\n", p1)
	//fmt.Printf("p1=%#v\n", childInt(p1)) // cannot convert p1 (type parent) to type childInt

	c1 := childInt{parent: p1, col: 1}
	fmt.Printf("c1=%#v\n", c1)
	c1.col = 2
	c1.i = 3
	c1.s = "xxx"
	c1.parent.col = 300
	fmt.Printf("c1=%#v\n", c1)
	//fmt.Printf("parent(c1)=%#v\n", parent(c1)) // cannot convert c1 (type childInt) to type parent
	fmt.Printf("c1.col=%#v\n", c1.col)

}
