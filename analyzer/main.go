package main

import (
	"github.com/tesujiro/goexp/analyzer/findcall" // my findcall
	//"golang.org/x/tools/go/analysis/passes/findcall"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(findcall.Analyzer)
}
