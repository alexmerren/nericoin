package main

import (
	"nericoin/internal/neri"
	"nericoin/internal/nerichain"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	Nerichain := nerichain.Nerichain{}
	TheOneAndOnlyNeri := neri.CreateGenesisBlock()
	Nerichain.AddNeri(TheOneAndOnlyNeri)
	spew.Dump(Nerichain)
}
