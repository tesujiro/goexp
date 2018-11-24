package main

import "fmt"

type key struct {
	key1 string
	key2 string
}

type list map[key][]int

var PriceList list

func init() {
	PriceList = make(list)
}

func main() {
	PriceList[key{key1: "AAA", key2: "Series1"}] = []int{1, 2, 3}
	PriceList[key{key1: "AAA", key2: "Series2"}] = []int{1, 2, 3}

	addPrice("AAA", "Series1", 100)
	addPrice("BBB", "Series1", 100)
	fmt.Printf("avg=%v\n", getPrice("AAA", "Series1"))
}

func addPrice(k1, k2 string, p int) {
	v, ok := PriceList[key{key1: k1, key2: k2}]
	if ok {
		v = append(v, p)
		fmt.Printf("v=%v\n", v)
		PriceList[key{key1: k1, key2: k2}] = v
	} else {
		PriceList[key{key1: k1, key2: k2}] = []int{p}
	}
}

func getPrice(k1, k2 string) int {
	prices, ok := PriceList[key{key1: k1, key2: k2}]
	if ok {
		total := 0
		for _, p := range prices {
			total += p
			fmt.Printf("price=%v\n", p)
		}
		fmt.Printf("total=%v\tlen=%v\n", total, len(prices))
		return total / len(prices)
	} else {
		return 0
	}
}
