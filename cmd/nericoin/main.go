package main

import (
	"fmt"
	"nericoin/internal/nerichain"
	"nericoin/internal/transaction"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	fmt.Println("Hello, world!")
	Nerichain := nerichain.NeriChain{}
	TheOneAndOnlyNeri := nerichain.CreateGenesisBlock()
	Nerichain = append(Nerichain, TheOneAndOnlyNeri)

	transactions := transaction.Transactions{}
	transactions.AddTransaction(transaction.Transaction{
		Ant:   "Jeremy",
		Onio:  "Joe",
		Value: 400,
	})
	transactions.AddTransaction(transaction.Transaction{
		Ant:   "Alex",
		Onio:  "Matt",
		Value: 420,
	})
	Nerichain.CreateNeri(transactions)
	spew.Dump(transactions)
}
