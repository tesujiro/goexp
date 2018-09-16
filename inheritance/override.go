package main

import "fmt"

type printable interface {
	print()
}

type parent struct {
	p_interface printable
}

func NewParent() *parent {
	p := &parent{}
	p.p_interface = p
	return p
}

func (p *parent) do() {
	p.p_interface.print()
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
	p.p_interface = p
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
