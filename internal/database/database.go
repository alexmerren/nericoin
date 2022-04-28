package database

import (
	"errors"
	"fmt"
	"nericoin/internal/neri"

	"github.com/boltdb/bolt"
)

type Database struct {
	db *bolt.DB
}

const (
	blocksBucket = "blocks"
	latestBucket = "latest"
	databaseFile = "nerichain.db"
)

func CreateDB() *Database {
	db, err := bolt.Open(databaseFile, 0600, nil)
	if err != nil {
		fmt.Println("Could not open or create database file")
		panic(err)
	}
	return &Database{db}
}

func (d *Database) CheckNerichainExists() bool {
	foundNerichain := false
	d.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		// if blockchain doesn't exist return false
		if b != nil {
			fmt.Println("Existing nerichain found")
			foundNerichain = true
		}
		return nil
	})
	return foundNerichain
}

func (d *Database) CreateBlockBucket() error {
	fmt.Println("Creating blocks bucket")
	err := d.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(blocksBucket))
		return err
	})
	return err
}

func (d *Database) InsertNeri(n *neri.Neri) error {

	fmt.Println("Inserting Neri into db...")
	serializedNeri, err := n.Serialize()
	if err != nil {
		return err
	}

	err = d.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err = b.Put([]byte(n.Hash), serializedNeri)
		if err != nil {
			return err
		}
		err = b.Put([]byte("l"), []byte(n.Hash))
		return err
	})
	return err
}

func (d *Database) GetLatestHash() (string, error) {
	var currentHash string
	err := d.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		currentHash = string(b.Get([]byte("l")))
		return nil
	})
	return currentHash, err
}

func (d *Database) GetNeri(hash string) (*neri.Neri, error) {
	var n *neri.Neri
	err := d.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		serializedNeri := b.Get([]byte(hash))
		n = neri.Deserialize(serializedNeri)
		if n == nil {
			return errors.New("could not deserialize current neri")
		}
		return nil
	})
	return n, err
}
