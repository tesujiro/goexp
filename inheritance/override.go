package main

import "fmt"

type overridable interface {
	print()
}

/*
type override struct {
	ov overridable
}
*/

type parent struct {
	override overridable
}

func NewParent() *parent {
	p := &parent{}
	//p.override = p
	return p
}

func (p *parent) do() {
	if p.override != nil {
		p.override.print()
	} else {
		p.print()
	}
}

func (p *parent) print() {
	fmt.Println("print parent")
}

// child
type child struct {
	parent
}

func NewChild() *child {
	p := &child{}
	p.override = p
	return p
}

func (c *child) print() {
	fmt.Println("print child")
}

//
func main() {
	p := NewParent()
	p.do()
	c := NewChild()
	c.do()
}
