package main

import "fmt"

func cartesian2(list [][]interface{}) [][]interface{} {
	var result [][]interface{}

	if 0 == len(list) {
		result = [][]interface{}{nil}
	} else {
		for _, e := range list[0] {
			for _, r := range cartesian2(list[1:]) {
				s := []interface{}{e}
				s = append(s, r...)
				result = append(result, s)
			}
		}
	}
	return result
}

func main() {
	var list [][]interface{}
	var result [][]interface{}

	list = [][]interface{}{{1, 2, 3}, {4, 5}, {6, 7}}
	result = cartesian2(list)
	fmt.Println("list:", list)
	fmt.Println("result:", result)

	list = [][]interface{}{{"xxx", "yyy"}, {"aaa", "bbb"}}
	result = cartesian2(list)
	fmt.Println("list:", list)
	fmt.Println("result:", result)
}
