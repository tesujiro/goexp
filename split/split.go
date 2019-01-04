package main

import "fmt"

func split_helper(advance int, token []byte, data []byte, pat []byte) (int, []byte) {
	if len(pat) == 0 || len(data) == 0 {
		return advance, token
	}
	if data[0] == pat[0] {
		return split_helper(advance+1, append(token, data[0]), data[1:], pat[1:])
	} else {
		//fmt.Println("NO MATCH")
		return split_helper(advance+1, append(token, data[0]), data[1:], pat)
	}
}

func split(data []byte, pat []byte) (int, []byte) {
	return split_helper(0, []byte{}, data, pat)
}

func main() {
	str := []byte("this_is_a_string")
	pat := []byte("is")

	fmt.Printf("str=%s\n", str)
	p := 0
	a, t := split(str[p:], pat)
	for a > 0 {
		fmt.Printf("a=%v token=%s\n", a, t)
		p += a
		a, t = split(str[p:], pat)
	}
	fmt.Printf("token=%s\n", t)
}
