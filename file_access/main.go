package main

import (
	"fmt"
	"io"
	"os"
)

const BlockSize = 1024

func initFile(fp *os.File, size int) error {
	blocks := size / BlockSize
	nullBlock := make([]byte, BlockSize)
	for i := 0; i < blocks; i++ {
		_, err := fp.Write(nullBlock)
		if err != nil {
			return err

		}
	}
	fmt.Printf("init %v bytes.\n", BlockSize*blocks)
	return nil
}

func main() {
	fp, err := os.Create("hello.block")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	err = initFile(fp, BlockSize*1024*10)
	if err != nil {
		fmt.Println(err)
		return
	}

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
		//fmt.Println(string(b))
	}
}
