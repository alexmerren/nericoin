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

const difficultRightNowLol = 1

// POLITE: Proof Of Life Inside The Element

type Polite struct {
	neri       *Neri
	difficulty int
}

func NewPolite(neri *Neri) *Polite {
	return &Polite{
		neri:       neri,
		difficulty: difficultRightNowLol,
	}
}

func (p *Polite) calculateHash() {
	data, _ := json.Marshal(p.neri.Data)
	toHash := p.neri.PreviousHash + p.neri.Timestamp.String() + string(data) + strconv.Itoa(p.difficulty) + strconv.Itoa(p.neri.TheElement)
	p.neri.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(toHash)))
}

func (p *Polite) Mine() {
	// Carries out the mining of the block by trying random nonces until the neri is verified

	attempt := 1
	fmt.Println("Mining...")
	for {
		//fmt.Println("Attemt " + strconv.Itoa(attempt))
		p.neri.TheElement = rand.Intn(math.MaxInt)

		if p.Verify() {
			fmt.Println(p.neri.Hash)
			fmt.Println("Mined in " + strconv.Itoa(attempt) + " attempts! Yayy!")
			break
		}
		attempt++
	}
}

func (p *Polite) Verify() bool {
	// Verify a successfully mined block by checking that the hash of the
	// block hash and the nonce is has the number of zeros defined by the
	// difficulty

	p.calculateHash()
	check_hash := p.neri.Hash

	return strings.HasPrefix(check_hash, strings.Repeat("0", p.difficulty))
}
