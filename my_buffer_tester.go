package main

import (
	"io"
	"io/ioutil"
	"os"
)

func main() {
	err := Double(os.Stdin, os.Stdout)
	if err != nil {
		panic(err)
	}
}

func Double(stdin io.Reader, stdout io.Writer) error {
	buf, err := ioutil.ReadAll(stdin)
	if err != nil {
		return err
	}

	stdout.Write(buf)
	stdout.Write(buf)

	return nil
}
