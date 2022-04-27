package nerichain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"nericoin/internal/transaction"
	"strconv"
	"time"
	"math/rand"
	"strings"
)

type Neri struct {
	hash         string
	previousHash string
	timestamp    time.Time
	data         transaction.Transactions
	difficulty   int
	nonce 		 int
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
	toHash := n.previousHash + n.timestamp.String() + string(data) + strconv.Itoa(n.difficulty) + strconv.Itoa(n.nonce)
	n.hash = fmt.Sprintf("%x", sha256.Sum256([]byte(toHash)))
}

func (n *Neri) Mine() {
	// Carries out the mining of the block by trying random nonces until the neri is verified

	attempt := 1
	fmt.Println("Mining...")
	for {
		//fmt.Println("Attemt " + strconv.Itoa(attempt))
		n.nonce = rand.Intn(100000)

		if n.Verify() == true {
			fmt.Println("Mined in " + strconv.Itoa(attempt) + " attempts! Yayy!")
			break
		}
		attempt++
	}
}

func (n *Neri) Verify() bool{
	// Verify a successfully mined block by checking that the hash of the
	// block hash and the nonce is has the number of zeros defined by the
	// difficulty

	n.calculateHash()
	check_hash := n.hash

	if strings.HasPrefix(check_hash, strings.Repeat("0", n.difficulty)) {
		return true
	}
	return false
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
	newNeri.Mine()
	fmt.Println("New Neri:")
	fmt.Println(newNeri)
	*n = append((*n), newNeri)
}
