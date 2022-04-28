package quickSum

import (
	"nericoin/internal/nerichain"
)

func WalletSum(qs *nerichain.Nerichain, addr string) int64 {
	// Work out the sum of a wallet

	neriIter := nerichain.NewNerichainIterator(qs)

	receivedTx := int64(0)
	sentTx := int64(0)
	for {
		neri := neriIter.GetNext()
		if neri == nil{
			break
		}
		if neri != nil {
			data := neri.Data
			//for _ , Tx := range data {
				if data.Ant == addr {
					//fmt.Println("Subtracting" + strconv.FormatInt(data.Value, 10))
					sentTx += data.Value
				}
				if data.Onio == addr {
					//fmt.Println("Adding" + strconv.FormatInt(data.Value, 10))
					receivedTx += data.Value
				}
			//}
		}
	}

	return receivedTx - sentTx
}