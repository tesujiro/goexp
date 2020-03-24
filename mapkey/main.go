package main

import (
	"fmt"
)

type t struct {
	key int
	val string
}

func main() {
	map_struct := make(map[t]string)

	map_struct[t{key: 1, val: "key1"}] = "val1"
	map_struct[t{key: 1, val: "key1"}] = "val2"
	fmt.Println(map_struct)

	map_struct_ptr := make(map[*t]string)
	map_struct_ptr[&t{key: 1, val: "key1"}] = "val1"
	map_struct_ptr[&t{key: 1, val: "key1"}] = "val2"
	fmt.Println(map_struct_ptr)

	//map_big := make(map[big.Int]string)
	//map_big[*big.NewInt(1)] = "val1"
	//fmt.Println(map_struct_ptr)

}
