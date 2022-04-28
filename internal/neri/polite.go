package neri

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

const currentDifficulty = 2

// PoLITE: Proof of Life Inside The Element

type Polite struct {
	neri       *Neri
	difficulty int
}

func NewPolite(neri *Neri) *Polite {
	return &Polite{
		neri:       neri,
		difficulty: currentDifficulty,
	}
}

func (p *Polite) calculateHash() {
	data, _ := json.Marshal(p.neri.Data)
	toHash := p.neri.PreviousHash + p.neri.Timestamp.String() + string(data) + strconv.Itoa(p.difficulty) + strconv.Itoa(p.neri.TheElement)
	p.neri.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(toHash)))
}

// Carries out the mining of the block by trying random nonces until the neri is verified
func (p *Polite) Mine() {
	attempt := 1
	fmt.Println(strings.Repeat("-", 10))
	fmt.Println("Mining...")
	for {
		p.neri.TheElement = rand.Intn(math.MaxInt)
		if p.Verify() {
			fmt.Println(p.neri.Hash)
			fmt.Println("Mined in " + strconv.Itoa(attempt) + " attempts! Yayy!")
			fmt.Println(strings.Repeat("-", 10))
			fmt.Println()
			break
		}
		attempt++
	}
}

// Verify a successfully mined block by checking that the hash of the
// block hash and the nonce is has the number of zeros defined by the
// difficulty
func (p *Polite) Verify() bool {
	p.calculateHash()
	check_hash := p.neri.Hash
	return strings.HasPrefix(check_hash, strings.Repeat("0", p.difficulty))
}
