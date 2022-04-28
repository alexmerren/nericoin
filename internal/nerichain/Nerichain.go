package nerichain

import (
	"fmt"
	"nericoin/internal/neri"

	"github.com/boltdb/bolt"
)

const (
	blocksBucket = "blocks"
	latestBucket = "latest"
	databaseFile = "nerichain.db"
)

type Nerichain struct {
	currentHash string
	database    *bolt.DB
}

func CreateNerichain() *Nerichain {
	var currentHash string

	// open BoltDB file, create if doesn't exist
	db, err := bolt.Open(databaseFile, 0600, nil)
	if err != nil {
		panic(err)
	}
	// check if blockchain exists in db
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		// if blockchain doesn't exist, then create
		if b == nil {
			fmt.Println("Creating new blockchain, no existing found...")
			theOneAndOnlyNeri := neri.CreateGenesisBlock()

			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				panic(err)
			}
			serializedNeri, _ := theOneAndOnlyNeri.Serialize()
			// add hash (key) and block (value)
			err = b.Put([]byte(theOneAndOnlyNeri.Hash), serializedNeri)
			if err != nil {
				panic(err)
			}
			// set l (key) to latest hash (value)
			err = b.Put([]byte("l"), []byte(theOneAndOnlyNeri.Hash))
			if err != nil {
				panic(err)
			}
			currentHash = theOneAndOnlyNeri.Hash
		} else {
			// if blockchain exists then get current hash
			currentHash = string(b.Get([]byte("l")))
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	return &Nerichain{currentHash, db}
}

func (n *Nerichain) AddNeri(data neri.Transaction) {
	var prevHash string

	// read only DB transaction to get latest block hash
	err := n.database.View(func(tx *bolt.Tx) error {
		// get blocks bucket
		b := tx.Bucket([]byte(blocksBucket))
		// get latest block hash
		prevHash = string(b.Get([]byte("l")))
		return nil
	})
	if err != nil {
		panic(err)
	}
	// create new neri (includes mining and setting hash)
	newNeri := neri.CreateNeri(prevHash, data)

	// open read-write DB transaction to add new block
	err = n.database.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		serializedNeri, _ := newNeri.Serialize()
		// add hash (key) and block (value)
		err := b.Put([]byte(newNeri.Hash), serializedNeri)
		if err != nil {
			panic(err)
		}
		// set l (key) to latest hash (value)
		err = b.Put([]byte("l"), []byte(newNeri.Hash))
		if err != nil {
			panic(err)
		}

		n.currentHash = newNeri.Hash

		return nil
	})
	if err != nil {
		panic(err)
	}
}

func (n *Nerichain) ViewNerichain() {
	iterator := NewNerichainIterator(n)
	for iterator.HasNext() {
		iterator.GetNext().String()
	}
}

// KILL ME
// KILL ME
// KILL ME

func (n *Nerichain) FindUnspentTransactions() string {
	return ""
}
