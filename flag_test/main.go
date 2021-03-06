package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	flags string
	flagt string
)

func main() {
	println(_main())
}

func _main() string {
	/* register flag name and shorthand name */
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.StringVar(&flags, "s", "", "string flag")
	f.StringVar(&flagt, "t", "", "string flag")

	f.Parse(os.Args[1:])
	/*
		for 0 < f.NArg() {
			f.Parse(f.Args()[1:])
		}
	*/

	return fmt.Sprintf("%v,%v,%v", flags, flagt, f.Args())
}
