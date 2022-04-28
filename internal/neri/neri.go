package neri

import (
	"time"
)

type Neri struct {
	Hash         string
	PreviousHash string
	Timestamp    time.Time
	Data         string
	TheElement   int
}

func CreateNeri(previousHash string, data string) *Neri {
	neri := &Neri{
		Timestamp:    time.Now(),
		PreviousHash: previousHash,
		Data:         data,
		TheElement:   0,  // The Element is intentially 0 for mining
		Hash:         "", // The Hash is calculated when mining the Neri
	}
	polite := NewPolite(neri)
	polite.Mine()
	return neri
}

func CreateGenesisBlock() *Neri {
	return CreateNeri("", "Praise our lord and saviour, the one and the only, Antonio Fabio Neri.")
}

func (n *Neri) Serialize() {}

func (n *Neri) Deserialize() {}
