package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	os.Exit(_main())
}

func _main() int {
	if len(os.Args) < 1 {
		fmt.Printf("error: no argument")
		return 1
	}

	for _, f := range os.Args[1:] {
		fmt.Printf("[%v]\n", f)
		fset := token.NewFileSet() // positions are relative to fset

		f, err := parser.ParseFile(fset, f, nil, parser.ParseComments)
		if err != nil {
			fmt.Println(err)
			return 1
		}

		ast.Print(nil, f)

		for _, i := range f.Imports {
			fmt.Printf("%s:import %v\n", fset.Position(i.Pos()), i.Path.Value)
		}
		for _, u := range f.Unresolved {
			fmt.Printf("%s:unresolved %v\n", fset.Position(u.Pos()), u.Name)
		}
		for _, c := range f.Comments {
			fmt.Printf("%s: %q\n", fset.Position(c.Pos()), c.Text())
		}
	}

	return 0
}

func parse(src string) (*ast.File, error) {

	fset := token.NewFileSet() // positions are relative to fset

	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return f, nil
}
