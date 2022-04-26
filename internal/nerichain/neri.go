package nerichain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"nericoin/internal/transaction"
	"strconv"
	"time"
)

type Neri struct {
	hash         string
	previousHash string
	timestamp    time.Time
	data         transaction.Transaction
	difficulty   int
}

func CreateGenesisBlock() *Neri {
	neri := &Neri{
		previousHash: "",
		timestamp:    time.Now(),
		data:         transaction.Transaction{},
		difficulty:   1,
	}
	neri.createHash()
	return neri
}

func (n *Neri) createHash() {
	data, _ := json.Marshal(n.data)
	toHash := n.previousHash + n.timestamp.String() + string(data) + strconv.Itoa(n.difficulty)
	n.hash = fmt.Sprintf("%x", sha256.Sum256([]byte(toHash)))
	fmt.Println("Hash of neri:\n" + n.hash)
}

func (n *NeriChain) CreateNeri(data transaction.Transaction) {
	prevNeri := (*n)[len(*n)-1]
	newNeri := &Neri{
		previousHash: prevNeri.hash,
		timestamp:    time.Now(),
		data:         data,
	}
	newNeri.mine()
	fmt.Println("New Neri:")
	fmt.Println(newNeri)
	*n = append((*n), newNeri)
}

func (n *Neri) mine() {
	fmt.Println("Mining...")
	(*n).createHash()
}
