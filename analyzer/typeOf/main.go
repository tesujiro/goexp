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
sub()
}
func sub() {
v := "abc"
println(v + "xyz")
}`

func main() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "main.go", src, 0)
	if err != nil {
		/* ؒٓ٦Ⳣ椚 */
		log.Fatal(err)
	}
	cfg := &types.Config{Importer: importer.Default()}
	info := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{},
	}
	_, err = cfg.Check("main", fset, []*ast.File{file}, info)
	if err != nil {
		/* ؒٓ٦Ⳣ椚 */
		log.Fatal(err)
	}
	ast.Inspect(file, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.BinaryExpr:
			pos := fset.Position(n.Pos())
			typ := info.TypeOf(n)
			fmt.Printf("type of %v is %v\n", pos, typ) // typ == types.Typ[types.Int]
			fmt.Printf("Type=%T\n", types.Typ)
			fmt.Printf("typ(%T)=%#v\n", typ, typ)
			fmt.Printf("typ=%#v\n", types.Typ[typ.(*types.Basic).Kind()])
		}
		return true
	})
}
