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
	data         transaction.Transactions
	difficulty   int
}

func CreateGenesisBlock() *Neri {
	neri := &Neri{
		previousHash: "",
		timestamp:    time.Now(),
		data:         transaction.Transactions{},
		difficulty:   1,
	}
	neri.calculateHash()
	return neri
}

func (n *Neri) calculateHash() {
	data, _ := json.Marshal(n.data)
	toHash := n.previousHash + n.timestamp.String() + string(data) + strconv.Itoa(n.difficulty)
	n.hash = fmt.Sprintf("%x", sha256.Sum256([]byte(toHash)))
	fmt.Println("Hash of neri:\n" + n.hash)
}

func (n *NeriChain) CreateNeri(data transaction.Transactions) {
	fmt.Println("Creating Neri...")
	prevNeri := (*n)[len(*n)-1]
	for _, t := range data {
		fmt.Println(t)
	}
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
	(*n).calculateHash()
}
