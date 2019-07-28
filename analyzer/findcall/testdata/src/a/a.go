// testdata/src/a/a.go
package main

import "fmt"

func main() {
	fmt.Println("hi") // want "call of Println"
	fmt.Print("hi")   // not a call of Println
}
