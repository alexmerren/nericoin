package transaction

type TransactionInput struct {
	TransactionID []byte
	Vout          int
	ScriptSig     string
}

func (t *TransactionInput) CanUnlockOutputWith(unlockingData string) bool {
	return t.ScriptSig == unlockingData
}
