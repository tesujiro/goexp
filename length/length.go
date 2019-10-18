package main

import (
	"fmt"
	"math/rand"

	"github.com/pkg/profile"
)

func main() {
	var l int
	N := 1000000
	number := make([]int, N)
	size := make([]int, N)
	fail := make([]int, N)
	for i := 0; i < N; i++ {
		size[i] = rand.Intn(1 << 32)
		number[i] = rand.Intn(size[i] + 1)
		fail[i] = rand.Intn(16)
	}
	fmt.Println("path1")

	defer profile.Start(profile.ProfilePath(".")).Stop()

	for i := 0; i < N; i++ {
		list := newBinSearchList(number[i])
		l = list.length()
	}
	fmt.Println("path2")
	for i := 0; i < N; i++ {
		list := newBinSearchListWithLimitedSize(number[i], size[i])
		l = list.length()
	}
	/*
		fmt.Println("path3")
		for i := 0; i < N; i++ {
			list := newSearchListWithLimitedFails(number[i], size[i], fail[i])
			l = list.length()
		}
	*/
	_ = l
}
