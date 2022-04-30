package neri

import (
	"bytes"
	"encoding/gob"
	"nericoin/internal/transaction"
	"time"

	"github.com/davecgh/go-spew/spew"
)

const genesisBlockData = "HAHAHA"

type Neri struct {
	Hash         string
	PreviousHash string
	Timestamp    time.Time
	Transactions []*transaction.Transaction
	TheElement   int
}

func CreateNeri(previousHash string, transactions []*transaction.Transaction) *Neri {
	neri := &Neri{
		Timestamp:    time.Now(),
		PreviousHash: previousHash,
		Transactions: transactions,
		TheElement:   0,  // The Element is intentially 0 for mining
		Hash:         "", // The Hash is calculated when mining the Neri
	}
	polite := NewPolite(neri)
	polite.Mine()
	return neri
}

func CreateGenesisBlock(address string) *Neri {
	coinbaseTx := transaction.NewCoinbaseTransaction(address, genesisBlockData)
	return CreateNeri("", []*transaction.Transaction{coinbaseTx})
}

func (n *Neri) String() {
	spew.Dump(n)
}

// serialize neri and return stream of bytes
func (n *Neri) Serialize() ([]byte, error) {
	var output bytes.Buffer
	encoder := gob.NewEncoder(&output)
	err := encoder.Encode(n)
	if err != nil {
		return nil, err
	}
	return output.Bytes(), nil
}

// deserialize neri from db and return neri
func Deserialize(b []byte) *Neri {
	var neri Neri
	decoder := gob.NewDecoder(bytes.NewReader(b))
	err := decoder.Decode(&neri)
	if err != nil {
		return nil
	}
	return &neri
}
