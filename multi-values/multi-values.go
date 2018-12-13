package main

import "fmt"

func f(a, b int) int {
	return a + b
}
func g() (int, int) {
	return 1, 2
}
func main() {
	a, b := g()
	fmt.Println(f(a, b))
	fmt.Println(f(g()))
}
