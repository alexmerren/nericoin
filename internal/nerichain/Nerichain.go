package nerichain

import (
	"encoding/hex"
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

// TODO Comment how this works
func (n *Nerichain) findUnspentTransactions(address string) []*transaction.Transaction {
	unspentTransactions := []*transaction.Transaction{}
	spentTransactions := make(map[string][]int)
	nerichainIterator := NewNerichainIterator(n)
	for {
		neri := nerichainIterator.GetNext()

		for _, transaction := range neri.Transactions {
			transactionID := hex.EncodeToString(transaction.ID)

			for outputIndex, output := range transaction.Vout {
				if spentTransactions[transactionID] != nil {
					for _, spentTransactionOutput := range spentTransactions[transactionID] {
						if spentTransactionOutput == outputIndex {
							continue
						}
					}
				}

				if output.CanBeUnlockedWith(address) {
					unspentTransactions = append(unspentTransactions, transaction)
				}
			}

			if !transaction.IsCoinbase() {
				for _, transactionInput := range transaction.Vin {
					if transactionInput.CanUnlockOutputWith(address) {
						inputTransactionID := hex.EncodeToString(transactionInput.TransactionID)
						spentTransactions[inputTransactionID] = append(spentTransactions[inputTransactionID], transactionInput.Vout)
					}
				}
			}
		}
		if len(neri.PreviousHash) == 0 {
			break
		}
	}
	return unspentTransactions
}

func (n *Nerichain) FindUTXOForAddress(address string) []transaction.TransactionOutput {
	unspentTransactions := n.findUnspentTransactions(address)
	UTXOs := make([]transaction.TransactionOutput, 0)

	for _, transaction := range unspentTransactions {
		for _, transactionOutput := range transaction.Vout {
			if transactionOutput.CanBeUnlockedWith(address) {
				UTXOs = append(UTXOs, transactionOutput)
			}
		}
	}

	return UTXOs
}
