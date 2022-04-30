package nerichain

import (
	"encoding/hex"
	"nericoin/internal/transaction"
)

// TODO Comment how this works
func (n *Nerichain) findUnspentTransactions(address string) []*transaction.Transaction {
	unspentTransactions := []*transaction.Transaction{}
	spentTransactions := make(map[string][]int)
	nerichainIterator := NewNerichainIterator(n)
	for {
		neri := nerichainIterator.GetNext()

		for _, transaction := range neri.Transactions {
			transactionID := hex.EncodeToString(transaction.ID)

		Outputs:
			for outputIndex, output := range transaction.Vout {
				if spentTransactions[transactionID] != nil {
					for _, spentTransactionOutput := range spentTransactions[transactionID] {
						if spentTransactionOutput == outputIndex {
							continue Outputs
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

// TODO Comment how this works
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

// TODO Comment how this works
func (n *Nerichain) FindSpendableOutputs(address string, amountRequested int) (int, map[string][]int) {
	unspentOutputs := make(map[string][]int)
	unspentTransactions := n.findUnspentTransactions(address)
	accumulated := 0

	for _, transaction := range unspentTransactions {
		transactionID := hex.EncodeToString(transaction.ID)
		for outputIndex, output := range transaction.Vout {
			if output.CanBeUnlockedWith(address) && accumulated < amountRequested {
				accumulated += output.Value
				unspentOutputs[transactionID] = append(unspentOutputs[transactionID], outputIndex)

				if accumulated >= amountRequested {
					break
				}
			}
		}
	}
	return accumulated, unspentOutputs
}

func (n *Nerichain) NewUTXOTransaction(from, to string, amount int) *transaction.Transaction {
	availableAmount, spendableOutputs := n.FindSpendableOutputs(from, amount)
	return transaction.NewUTXOTransaction(from, to, amount, availableAmount, spendableOutputs)
}

func (n *Nerichain) NewCoinbaseTransaction(to, data string) *transaction.Transaction {
	return transaction.NewCoinbaseTransaction(to, data)
}
