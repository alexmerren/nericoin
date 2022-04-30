package transaction

import "fmt"

const reward = 10

func NewCoinbaseTransaction(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}
	input := TransactionInput{[]byte{}, -1, data}
	output := TransactionOutput{reward, to}
	transaction := &Transaction{nil, []TransactionInput{input}, []TransactionOutput{output}}
	transaction.SetID()
	return transaction
}

func (t *Transaction) IsCoinbase() bool {
	return len(t.Vin) == 1 && len(t.Vin[0].TransactionID) == 0 && t.Vin[0].Vout == -1
}
