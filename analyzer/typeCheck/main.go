package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

const src = `package main
func main() {
v := 100
println(v + 1)
}`

func main() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "main.go", src, 0)
	if err != nil {
		/* ؒٓ */
		log.Fatal(err)
	}
	cfg := &types.Config{Importer: importer.Default()}
	info := &types.Info{
		Defs: map[*ast.Ident]types.Object{},
	}
	pkg, err := cfg.Check("main", fset, []*ast.File{file}, info)
	if err != nil {
		/* ؒٓ٦Ⳣ椚 */
		log.Fatal(err)
	}
	fmt.Println("package is", pkg.Path())
	for ident, obj := range info.Defs {
		fmt.Println(ident.Name, obj)
	}
}
