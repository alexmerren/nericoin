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

func (n *Neri) createHash() string{
	data, _ := json.Marshal(n.data)
	toHash := n.previousHash + n.timestamp.String() + string(data) + string(n.difficulty)
	n.hash = fmt.Sprintf("%x", sha256.Sum256([]byte(toHash)))
	return n.hash
}

func (n *Neri) VerifyNeri() bool{
	// Verify a successfully mined block by checking that the hash of the
	// block hash and the nonce is has the number of zeros defined by the
	// difficulty

	fmt.Println("Verifying Neri")
	verified_status := false

	//Perform hash function on the combined block_hash and nonce
	check_hash := n.createHash()
	//check_hash = "0x00000000454"
	fmt.Println("Hash: " + check_hash)

	// Validate the number of leading zeros vs difficulty
	num_zeros := countZeros(check_hash)
	//fmt.Println("Num Zeros: " + num_zeros)

	if num_zeros >= n.difficulty{
		verified_status = true
	}
	return verified_status
}

func countZeros(hash string) int{
	// Split the hash up into single runes and compare each rune to the ascii
	// for 0 (48) and if yes then increment the total number of zeros. Else
	// break the for loop and return num_zeros.
	digits := []rune(hash)
	num_zeros := 0
	for i := 0; i < len(digits); i++ {
		if digits[i] == 48 {
			num_zeros++
		} else {
			break
		}
	}
	return num_zeros
}
