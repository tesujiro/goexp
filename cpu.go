package main

import "runtime"
import "fmt"

func main() {
	fmt.Println("%d\n", runtime.NumCPU())
}
