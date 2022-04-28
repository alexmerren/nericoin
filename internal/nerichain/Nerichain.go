package nerichain

import (
	"nericoin/internal/database"
	"nericoin/internal/neri"
)

type Nerichain struct {
	currentHash string
	db          *database.Database
}

func CreateNerichain() *Nerichain {
	var currentHash string
	db := database.CreateDB()
	if db.CheckNerichainExists() {
		currentHash, _ = db.GetLatestHash()
	} else {
		db.CreateBlockBucket()
		theOneAndOnlyNeri := neri.CreateGenesisBlock()
		db.InsertNeri(theOneAndOnlyNeri)
		currentHash, _ = db.GetLatestHash()
	}
	return &Nerichain{currentHash, db}
}

func (n *Nerichain) AddNeri(data neri.Transaction) {
	prevHash, _ := n.db.GetLatestHash()

	// create new neri (includes mining and setting hash)
	newNeri := neri.CreateNeri(prevHash, data)
	n.db.InsertNeri(newNeri)
	n.currentHash = newNeri.Hash
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
