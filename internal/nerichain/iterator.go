package nerichain

import (
	"nericoin/internal/neri"
)

type NerichainIterator struct {
	currentHash string
	nerichain   *Nerichain
}

func NewNerichainIterator(n *Nerichain) *NerichainIterator {
	return &NerichainIterator{
		currentHash: n.currentHash,
		nerichain:   n,
	}
}

// View the database and get the neri with the currentHash in the
// iterator. Return that neri.
func (i *NerichainIterator) GetNext() *neri.Neri {
	if !i.HasNext() {
		return nil
	}
	nextNeri, _ := i.nerichain.db.GetNeri(i.currentHash)
	i.currentHash = nextNeri.PreviousHash

	return nextNeri
}

// Get neri with current hash from database and get the previous hash
// and the set the iterators current hash as the next (previous) Neri
// in the chain.  If there is no Neri attached to a hash, there is
// either an error or you've reached the start.
func (i *NerichainIterator) HasNext() bool {
	n, _ := i.nerichain.db.GetNeri(i.currentHash)
	return n != nil
}
