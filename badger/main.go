package main

import (
	"fmt"
	"log"

	badger "github.com/dgraph-io/badger/v2"
)

func main() {
	//opt := badger.DefaultOptions("testdb")
	opt := badger.DefaultOptions("").WithInMemory(true)
	db, err := badger.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Start a writable transaction.
	txn := db.NewTransaction(true)
	defer txn.Discard()

	key := "answer"
	value := "42"
	// Use the transaction...
	err = txn.Set([]byte(key), []byte(value))
	if err != nil {
		fmt.Printf("Set err:%v\n", err)
		return
	}

	// Commit the transaction and check for error.
	if err := txn.Commit(); err != nil {
		fmt.Printf("Commit err:%v\n", err)
		return
	}

	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			fmt.Printf("The answer is: %s\n", val)
			return nil
		})

		return nil
	})
	if err != nil {
		fmt.Printf("View err:%v\n", err)
		return
	}
}
