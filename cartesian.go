package main

import "fmt"

func cartesian(n int, list [][]interface{}) [][]interface{} {
	var result [][]interface{}

	if n == len(list) {
		result = [][]interface{}{nil}
	} else {
		for _, e := range list[n] {
			for _, r := range cartesian(n+1, list) {
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
	result = cartesian(0, list)
	fmt.Println("list:", list)
	fmt.Println("result:", result)

	list = [][]interface{}{{"xxx", "yyy"}, {"aaa", "bbb"}}
	result = cartesian(0, list)
	fmt.Println("list:", list)
	fmt.Println("result:", result)
}
