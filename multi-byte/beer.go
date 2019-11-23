package main

import (
	"fmt"
	"unicode"
)

func main() {
	beer_en := "beer"
	beer_jp := "ビール"
	beer_emoji := "🍺"
	beer_str := `beer!ビール!🍺!!!`
	//🍺:="🍺" // error

	fmt.Printf("%s\n", beer_en)
	fmt.Printf("%s\n", beer_jp)
	fmt.Printf("%s\n", beer_emoji)
	fmt.Printf("%s\n", beer_str)
	//fmt.Printf("%s\n", 🍺)

	for i, r := range beer_str {
		fmt.Printf("%d:\t%c\tIsLetter:%v\n", i, r, unicode.IsLetter(r))
	}
}
