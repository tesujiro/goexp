package main

import (
	"fmt"
	"io"
	"os"
)

const BlockSize = 1024

func main() {
	fp, err := os.Create("hello.binary")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	null := make([]byte, BlockSize)

	for i := range null {
		null[i] = '.'
	}
	fp.Write(null)

	numbers := []byte("0123456789")
	fp.Seek(256, os.SEEK_SET)
	fp.Write(numbers)

	b := make([]byte, BlockSize)
	fp.Seek(0, 0)
	for {
		_, err := fp.Read(b)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
				return
			}
			break
		}
		fmt.Println(string(b))
	}
}
