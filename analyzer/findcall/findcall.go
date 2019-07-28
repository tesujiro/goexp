package findcall

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

const Doc = `find calls to a particular function
The findcall analysis reports calls to functions or methods
of a particular name.`

var Analyzer = &analysis.Analyzer{
	Name:             "findcall",
	Doc:              Doc,
	Run:              run,
	RunDespiteErrors: true,
	FactTypes:        []analysis.Fact{new(foundFact)},
}

var name string // -name flag

func init() {
	Analyzer.Flags.StringVar(&name, "name", name, "name of the function to find")
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(n ast.Node) bool {
			if call, ok := n.(*ast.CallExpr); ok {
				var id *ast.Ident
				switch fun := call.Fun.(type) {
				case *ast.Ident:
					id = fun
					//fmt.Printf("id=%v\n", id)
				case *ast.SelectorExpr:
					id = fun.Sel
					//fmt.Printf("selector id=%v\n", id)
				}
				if id != nil && !pass.TypesInfo.Types[id].IsType() && id.Name == name {
					pass.Reportf(call.Lparen, "call of %s(...)", id.Name)
				}
			}
			return true
		})
	}

	// Export a fact for each matching function.
	//
	// These facts are produced only to test the testing
	// infrastructure in the analysistest package.
	// They are not consumed by the findcall Analyzer
	// itself, as would happen in a more realistic example.
	for _, f := range pass.Files {
		for _, decl := range f.Decls {
			if decl, ok := decl.(*ast.FuncDecl); ok && decl.Name.Name == name {
				if obj, ok := pass.TypesInfo.Defs[decl.Name].(*types.Func); ok {
					pass.ExportObjectFact(obj, new(foundFact))
				}
			}
		}
	}

	return nil, nil
}

type foundFact struct{}

func (*foundFact) String() string { return "found" }
func (*foundFact) AFact()         {}
