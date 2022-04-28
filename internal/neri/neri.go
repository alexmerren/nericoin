package neri

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Neri struct {
	Hash         string
	PreviousHash string
	Timestamp    time.Time
	Data         Transaction
	TheElement   int
}

type Transaction struct {
	Ant   string
	Onio  string
	Value int64
}

func CreateNeri(previousHash string, data Transaction) *Neri {
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

	return CreateNeri("0", Transaction{
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

func (n *Neri) Serialize() ([]byte, error) {
	// serialize neri and return stream of bytes
	var output bytes.Buffer
	encoder := gob.NewEncoder(&output)
	err := encoder.Encode(n)
	if err != nil {
		return nil, err
	}
	return output.Bytes(), nil
}

func Deserialize(b []byte) *Neri {
	// deserialize neri from db and return neri
	var neri Neri
	decoder := gob.NewDecoder(bytes.NewReader(b))
	err := decoder.Decode(&neri)
	if err != nil {
		return nil
	}
	return &neri
}
