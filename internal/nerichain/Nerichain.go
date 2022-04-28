package nerichain

import (
	"nericoin/internal/neri"

	"github.com/boltdb/bolt"
)

type Nerichain struct {
	currentHash string
	database    *bolt.DB
}

func CreateNerichain() *Nerichain {
	return &Nerichain{}
}

func (n *Nerichain) AddNeri(neri *neri.Neri) {}

// KILL ME
// KILL ME
// KILL ME

func (n *Nerichain) FindUnspentTransactions() string {
	return ""
}
