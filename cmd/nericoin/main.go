package main

import (
	"fmt"
	"nericoin/internal/nerichain"
	"nericoin/internal/transaction"
)

func main() {
	fmt.Println("Hello, world!")
	Nerichain := nerichain.NeriChain{}
	TheOneAndOnlyNeri := nerichain.CreateGenesisBlock()
	Nerichain = append(Nerichain, TheOneAndOnlyNeri)
	Nerichain.CreateNeri(transaction.Transaction{Ant: "Antonio", Onio: "Sam H", Value: 420})
	transactions := transaction.Transactions{}
	transactions.AddTransaction(transaction.Transaction{
		Ant:   "Jeremy",
		Onio:  "Joe",
		Value: 400,
	})
	transactions.AddTransaction(transaction.Transaction{
		Ant:   "Alex",
		Onio:  "Matt",
		Value: 69,
	})
}
