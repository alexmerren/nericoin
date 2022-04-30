package nerichain

import (
	"nericoin/internal/database"
	"nericoin/internal/neri"
	"nericoin/internal/transaction"
)

type Nerichain struct {
	currentHash string
	db          *database.Database
}

func CreateNerichain(address string) *Nerichain {
	var currentHash string
	db := database.CreateDB()
	if db.CheckNerichainExists() {
		currentHash, _ = db.GetLatestHash()
	} else {
		db.CreateBlockBucket()
		theOneAndOnlyNeri := neri.CreateGenesisBlock(address)
		db.InsertNeri(theOneAndOnlyNeri)
		currentHash, _ = db.GetLatestHash()
	}
	return &Nerichain{currentHash, db}
}

// create new neri (includes mining and setting hash)
func (n *Nerichain) AddNeri(transactions []*transaction.Transaction) {
	prevHash, _ := n.db.GetLatestHash()
	newNeri := neri.CreateNeri(prevHash, transactions)
	n.db.InsertNeri(newNeri)
	n.currentHash = newNeri.Hash
}

func (n *Nerichain) ViewNerichain() {
	iterator := NewNerichainIterator(n)
	for iterator.HasNext() {
		iterator.GetNext().String()
	}
}
