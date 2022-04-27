package nerichain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
	"math/rand"
	"strconv"
	"strings"
)

type Neri struct {
	hash         string
	previousHash string
	timestamp    time.Time
	data         map[string]interface{}
	difficulty   int
	nonce 		 int
}

func CreateGenesisBlock() *Neri {
	neri := &Neri{
		previousHash: "",
		timestamp:    time.Now(),
		data:         map[string]interface{}{},
		difficulty:   2,
	}
	neri.createHash()
	return neri
}

func (n *Neri) createHash() {
	data, _ := json.Marshal(n.data)
	toHash := n.previousHash + n.timestamp.String() + string(data) + string(n.difficulty) + string(n.nonce)
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

	n.createHash()
	check_hash := n.hash

	if strings.HasPrefix(check_hash, strings.Repeat("0", n.difficulty)) {
		return true
	}
	return false
}