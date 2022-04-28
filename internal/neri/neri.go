package neri

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"nericoin/internal/transaction"
	"strconv"
	"strings"
	"time"
)

type Neri struct {
	Hash         string
	PreviousHash string
	Timestamp    time.Time
	Data         transaction.Transaction
	TheElement   int
}

func CreateNeri(previousHash string, data transaction.Transaction) *Neri {
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
	return CreateNeri("0", transaction.Transaction{
		Ant:   "Antonio",
		Onio:  "Nimble",
		Value: 1000000000,
	})
}

func (n *Neri) String() {
	l := len(n.Hash) + 20
	fmt.Println(strings.Repeat("-", l))
	hashPrint := "| Hash: " + n.Hash
	prevHashPrint := "| Previous Hash: " + n.PreviousHash
	antPrint := "| Ant: " + n.Data.Ant
	onioPrint := "| Onio: " + n.Data.Onio
	valPrint := "| Value: " + strconv.FormatInt(n.Data.Value, 10)
	timestampPrint := "| Timestamp: " + n.Timestamp.String()
	elementPrint := "| The Element: " + strconv.Itoa(n.TheElement)
	fmt.Println(hashPrint + strings.Repeat(" ", l-len(hashPrint)) + "|")
	fmt.Println(prevHashPrint + strings.Repeat(" ", l-len(prevHashPrint)) + "|")
	fmt.Println(antPrint + strings.Repeat(" ", l-len(antPrint)) + "|")
	fmt.Println(onioPrint + strings.Repeat(" ", l-len(onioPrint)) + "|")
	fmt.Println(valPrint + strings.Repeat(" ", l-len(valPrint)) + "|")
	fmt.Println(timestampPrint + strings.Repeat(" ", l-len(timestampPrint)) + "|")
	fmt.Println(elementPrint + strings.Repeat(" ", l-len(elementPrint)) + "|")
	fmt.Println(strings.Repeat("-", l))
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
