package main

import (
	"fmt"
	"go/parser"
	"go/token"

	"golang.org/x/tools/go/ast/astutil"
)

const src = `package main
var v = 100
func main() {
fmt.Println(v+1)
}`

func main() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "main.go", src, 0)
	if err != nil {
		/* ؒٓ*/
	}
	fset.Iterate(func(f *token.File) bool {
		if f.Name() == "main.go" {
			pos := token.Pos(f.LineStart(2) + 4) // var "v" = 100
			path, exact := astutil.PathEnclosingInterval(file, pos, pos)
			if exact {
				for _, n := range path {
					fmt.Printf("%T\n", n)
				}
			}
		}
		return true
	})
}
