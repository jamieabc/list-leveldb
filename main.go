package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/fatih/color"
	ldb "github.com/syndtr/goleveldb/leveldb"
)

func main() {
	if 3 != len(os.Args) {
		fmt.Printf("Usage: list-leveldb path-to-db key")
		return
	}

	dbPath := os.Args[1]
	key := os.Args[2]
	cyan := color.New(color.FgCyan).PrintfFunc()
	yellow := color.New(color.FgYellow).PrintfFunc()

	db, err := ldb.OpenFile(dbPath, nil)
	defer db.Close()

	if nil != err {
		fmt.Printf("Error: open level db %s: %s\n", dbPath, err)
		return
	}

	switch key {
	case "list":
		iter := db.NewIterator(nil, nil)
		fmt.Printf("list database\n")
		for iter.Next() {
			key := iter.Key()
			value := iter.Value()
			cyan("key: %x\n", key)
			yellow("value: %x\n", value)
		}
	default:
		hexKey, _ := hex.DecodeString(key)
		data, err := db.Get(hexKey, nil)
		if nil != err {
			yellow("Error: get key %x: %s\n", hexKey, err)
			return
		}
		cyan("key: %x\n", hexKey)
		yellow("value: %x\n", data)
	}

}
