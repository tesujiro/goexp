package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func readCsv(in io.Reader) {
	//input := bufio.NewScanner(in)
	input := csv.NewReader(in)
	for {
		record, err := input.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Println(record)
	}
}

func main() {
	readCsv(os.Stdin)
}
