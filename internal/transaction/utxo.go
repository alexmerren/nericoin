package transaction

import (
	"encoding/hex"
	"log"
)

func NewUTXOTransaction(
	from, to string,
	amountRequested, availableAmount int,
	spendableOutputs map[string][]int,
) *Transaction {
	var inputs []TransactionInput
	var outputs []TransactionOutput

	if amountRequested > availableAmount {
		log.Fatal("Not enough funds")
	}

	for id, outputs := range spendableOutputs {
		transactionID, _ := hex.DecodeString(id)
		for _, output := range outputs {
			input := TransactionInput{transactionID, output, from}
			inputs = append(inputs, input)
		}
	}

	outputs = append(outputs, TransactionOutput{amountRequested, to})
	if availableAmount > amountRequested {
		outputs = append(outputs, TransactionOutput{availableAmount - amountRequested, from})
	}

	transaction := &Transaction{nil, inputs, outputs}
	transaction.SetID()
	return transaction
}
