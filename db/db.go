package db

/*
1. create a transaction
Transaction
how do create bucket?

Bolot allows only one read-write transaction at a time but allows as many read-only transactions as you want at a time. Each transaction has a consistent view of the data as it existed when the transaction started.

individual trasnactions and all objects created form them. are not thread safe. to work with data in multiple goroutines you must start a transaction for each one or use locking to ensure only one goroutine accesses a trasaction at a time. creating trasaction from the (db) is thread safe
*/

import (
	"fmt"

	"github.com/DongSeonYoo/go-coin/utils"
	"github.com/boltdb/bolt"
)

const (
	dbName       = "blockchain.db"
	dataBucket   = "data"
	blocksBucket = "blocks"
	checkpoint   = "checkpoint"
)

var db *bolt.DB

func DB() *bolt.DB {
	if db == nil {
		// init database
		dbPointer, err := bolt.Open(dbName, 0600, nil)
		db = dbPointer
		utils.HandleErr(err)
		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_, err = tx.CreateBucketIfNotExists([]byte(blocksBucket))
			return err
		})
		utils.HandleErr(err)
	}

	return db
}

func SaveBlock(hash string, data []byte) {
	err := DB().Update(func(tx *bolt.Tx) error {
		blockBucket := tx.Bucket([]byte(blocksBucket))
		err := blockBucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleErr(err)
}

func SaveBlockChain(data []byte) {
	err := DB().Update(func(tx *bolt.Tx) error {
		dataBucket := tx.Bucket([]byte(dataBucket))
		err := dataBucket.Put([]byte(checkpoint), data)
		return err
	})
	utils.HandleErr(err)
}

func CheckPoint() []byte {
	var data []byte
	fmt.Println("staring checkpoint")
	DB().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(checkpoint))

		return nil
	})

	return data
}
