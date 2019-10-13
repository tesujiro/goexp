package main

import (
	"math/rand"

	"github.com/pkg/profile"
)

func main() {
	var l int
	N := 10000000
	defer profile.Start(profile.ProfilePath(".")).Stop()
	for i := 0; i < N; i++ {
		rnd := rand.Intn(N)
		s := newBinSearchListWithLimitedSize(rnd, rnd+1)
		l = s.length()
	}
	_ = l
}
