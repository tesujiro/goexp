package main

import "fmt"

type Parent interface{}

type Child struct {
	number int
}

func f1(p Parent) {
	c := p.(Child)
	fmt.Printf("c=%#v\n", c)
	p = Child{number: 3}
}

func f2(p *Parent) {
	c := (*p).(Child)
	fmt.Printf("c=%#v\n", c)
	*p = Child{number: 3}
}

func f3(p Parent) {
	c := p.(*Child)
	fmt.Printf("c=%#v\n", c)
	*c = Child{number: 3}
}

func main() {
	c1 := Child{number: 1}
	//c2 := Child{number: 2}

	// BAD
	fmt.Printf("c1=%#v\n", c1)
	f1(c1)
	fmt.Printf("c1=%#v\n", c1)

	// NOT GOOD
	c1p := Parent(c1)
	fmt.Printf("c1p=%#v\n", c1p)
	f2(&c1p)
	fmt.Printf("c1p=%#v\n", c1p)

	// GOOD!
	fmt.Printf("c1=%#v\n", c1)
	f3(&c1)
	fmt.Printf("c1=%#v\n", c1)

}
