package nerichain

import (
	"errors"
	"nericoin/internal/neri"

	"github.com/boltdb/bolt"
)

type NerichainIterator struct {
	currentHash string
	nerichain   *Nerichain
}

func NewNerichainIterator(n *Nerichain) *NerichainIterator {
	return &NerichainIterator{
		nerichain: n,
	}
}

// View the database and get the serialized neri with the currentHash in the
// iterator.  Deserialize the neri and return that neri.
func (i *NerichainIterator) GetNext() *neri.Neri {
	if !i.HasNext() {
		return nil
	}

	var nextNeri *neri.Neri
	err := i.nerichain.database.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		serializedNeri := bucket.Get([]byte(i.currentHash))
		neri := neri.Deserialize(serializedNeri)
		if neri == nil {
			return errors.New("")
		}
		nextNeri = neri
		return nil
	})
	if err != nil {
		return nil
	}
	return nextNeri
}

// View the database and find the serialized Neri that has the currentHash.
// Then, Deserialize the Neri to get the previous hash, and the set the
// iterators current hash as the next (previous) Neri in the chain.  If there
// is no Neri attached to a hash, there is either an error or you've reached
// the start.
func (i *NerichainIterator) HasNext() bool {
	err := i.nerichain.database.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		serializedNeri := bucket.Get([]byte(i.currentHash))
		if serializedNeri != nil {
			return errors.New("Could not find the previous hash")
		}
		neri := neri.Deserialize(serializedNeri)
		if neri == nil {
			return errors.New("")
		}
		i.currentHash = neri.PreviousHash
		return nil
	})
	if err != nil {
		return false
	}
	return true
}
