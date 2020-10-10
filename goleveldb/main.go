package main

import (
	"bytes"
	"fmt"
	"math/rand"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func main() {
	db, err := leveldb.OpenFile("testdata", nil)
	if err != nil {
		fmt.Printf("Open failed: %v", err)
	}
	defer db.Close()

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

		db.Put([]byte(key), []byte(val), nil)
		if err != nil {
			fmt.Printf("Put(%q): %v\n", key, err)
		}

		g, err := db.Get([]byte(key), nil)
		if err != nil {
			fmt.Printf("Get(%q): %v\n", key, err)
		}

		if bytes.Compare(g, []byte(val)) != 0 {
			fmt.Printf("Get(%q): result error want:%v got:%v\n", key, val, g)
		}
	}

	// Range: ALL
	//iter := db.NewIterator(nil, nil)
	// Range: random
	start := keys[rand.Intn(len(keys)/2)]
	limit := keys[len(keys)/2+rand.Intn(len(keys)/2)]
	iter := db.NewIterator(&util.Range{Start: []byte(start), Limit: []byte(limit)}, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("key:%s\tvalue:%s\n", key, value)
	}
	iter.Release()

}
