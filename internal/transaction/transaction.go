package transaction

import "fmt"

type Transaction struct {
	Ant   string
	Onio  string
	Value float64
}

type Transactions []*Transaction

func (t *Transactions) AddTransaction(transaction Transaction) {
	fmt.Println("Send transaction to each miner...")
	fmt.Println(transaction)
	*t = append((*t), &transaction)
}
