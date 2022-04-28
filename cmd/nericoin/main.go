package main

import (
	"nericoin/internal/nerichain"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	Nerichain := nerichain.CreateNerichain()
	Nerichain.AddNeri("Daddy")
	Nerichain.AddNeri("Mummy")
	Nerichain.AddNeri("Uncle")
	spew.Dump(Nerichain)
}
