package findcall

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func init() {
	Analyzer.Flags.Set("name", "Println")
}

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "a")
}
