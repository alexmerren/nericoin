package transaction

type TransactionOutput struct {
	Value        int
	ScriptPubKey string
}

func (t *TransactionOutput) CanBeUnlockedWith(unlockingData string) bool {
	return t.ScriptPubKey == unlockingData
}
