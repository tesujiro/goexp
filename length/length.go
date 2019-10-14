package main

import (
	"fmt"
)

func main() {
	list := newSearchListWithLimitedFails(129, 256, 3)
	fmt.Println(list.length())
	/*
		var l int
		N := 10000000
		defer profile.Start(profile.ProfilePath(".")).Stop()
			for i := 0; i < N; i++ {
				rnd := rand.Intn(N)
				s := newBinSearchListWithLimitedSize(rnd, rnd+1)
				l = s.length()
			}
			_ = l
	*/
}
