package neri

import (
	"crypto/sha256"
	"fmt"
	"log"
	"math"
	"math/rand"
	"nericoin/internal/transaction"
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
	toHash := p.neri.PreviousHash +
		p.neri.Timestamp.String() +
		hashTransactions(p.neri.Transactions) +
		strconv.Itoa(p.difficulty) +
		strconv.Itoa(p.neri.TheElement)
	p.neri.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(toHash)))
}

// Carries out the mining of the block by trying random nonces until the neri is verified
func (p *Polite) Mine() {
	attempt := 1
	for {
		p.neri.TheElement = rand.Intn(math.MaxInt)
		if p.Verify() {
			log.Printf("\nMined %s in %d attempts\n", p.neri.Hash, attempt)
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
	return strings.HasPrefix(p.neri.Hash, strings.Repeat("0", p.difficulty))
}

func hashTransactions(transactions []*transaction.Transaction) string {
	// TODO: Actually implement this
	return ""
}
