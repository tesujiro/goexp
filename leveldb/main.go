package main

import (
	"bytes"
	"fmt"

	"github.com/golang/leveldb"
	"github.com/golang/leveldb/db"
)

func main() {
	/*
		d, err := leveldb.Open("", &db.Options{ FileSystem: memfs.New(), })
	*/
	d, err := leveldb.Open("testdata", nil)
	if err != nil {
		fmt.Printf("Open failed: %v", err)
	}

	size := 10000
	keys := make([]string, size)
	vals := make([]string, size)
	for i := 0; i < size; i++ {
		keys[i] = fmt.Sprintf("key%.5d", i)
		vals[i] = fmt.Sprintf("val%.5d", i)
	}

	for i := 0; i < size; i++ {
		key := keys[i]
		val := vals[i]

		d.Set([]byte(key), []byte(val), nil)
		if err != nil {
			fmt.Printf("Set(%q): %v\n", key, err)
		}

		g, err := d.Get([]byte(key), nil)
		if err != nil && err != db.ErrNotFound {
			fmt.Printf("Get(%q): %v\n", key, err)
		}

		if bytes.Compare(g, []byte(val)) != 0 {
			fmt.Printf("Get(%q): result error want:%v got:%v\n", key, val, g)
		}
	}

	//time.Sleep(5 * time.Second)
	if err := d.Close(); err != nil {
		fmt.Printf("Close failed: %v", err)
	}
}
