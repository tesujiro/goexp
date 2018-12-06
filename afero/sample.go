package main

import (
	"fmt"
	"os"

	"github.com/spf13/afero"
)

func writeAndCheck(fs afero.Fs) {
	// create test files and directories
	fs.MkdirAll("./src/a", 0755)
	afero.WriteFile(fs, "./src/a/b", []byte("file b"), 0644)
	afero.WriteFile(fs, "./src/c", []byte("file c"), 0644)
	afero.WriteFile(fs, "./src/あああ", []byte("file あああ"), 0644)
	name := "./src/c"
	_, err := fs.Stat(name)
	if os.IsNotExist(err) {
		fmt.Printf("file \"%s\" does not exist.\n", name)
	}
	if b, err := afero.ReadFile(fs, name); err != nil {
		fmt.Printf("file \"%s\" read error :%v\n", err)
	} else {
		fmt.Printf("%v\n", string(b))
	}
	if err := fs.RemoveAll("./src"); err != nil {
		fmt.Printf("Remove All error :%v\n", err)
	}
}

func main() {
	appFS := afero.NewMemMapFs()
	writeAndCheck(appFS)
	osFS := afero.NewOsFs()
	writeAndCheck(osFS)
}
