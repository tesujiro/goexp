package main

import (
	"fmt"
	"sort"
)

func main() {

	m := map[string]int{"Alice": 23, "Eve": 2, "Bob": 25}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}

	fmt.Println(keys)
}
