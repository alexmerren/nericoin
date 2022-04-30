package transaction

type Transaction struct {
	ID   []byte
	Vin  []TransactionInput
	Vout []TransactionOutput
}

func (t *Transaction) SetID() {}
