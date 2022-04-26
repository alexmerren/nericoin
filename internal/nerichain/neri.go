package nerichain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Neri struct {
	hash         string
	previousHash string
	timestamp    time.Time
	data         map[string]interface{}
	difficulty   int
}

func CreateGenesisBlock() *Neri {
	neri := &Neri{
		previousHash: "",
		timestamp:    time.Now(),
		data:         map[string]interface{}{},
		difficulty:   1,
	}
	neri.createHash()
	return neri
}

func (n *Neri) createHash() {
	data, _ := json.Marshal(n.data)
	toHash := n.previousHash + n.timestamp.String() + string(data) + string(n.difficulty)
	n.hash = fmt.Sprintf("%x", sha256.Sum256([]byte(toHash)))
	fmt.Println("Hash of neri:\n" + n.hash)
}
