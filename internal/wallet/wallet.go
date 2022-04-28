package wallet

import (
	"nericoin/internal/nerichain"
)

// Work out the sum of a wallet
func WalletSum(qs *nerichain.Nerichain, addr string) int64 {
	neriIter := nerichain.NewNerichainIterator(qs)
	receivedTx := int64(0)
	sentTx := int64(0)
	for {
		neri := neriIter.GetNext()
		if neri == nil {
			break
		}
		if neri != nil {
			data := neri.Data
			if data.Ant == addr {
				sentTx += data.Value
			}
			if data.Onio == addr {
				receivedTx += data.Value
			}
		}
	}

	return receivedTx - sentTx
}
