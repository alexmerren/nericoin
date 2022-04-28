package main

import (
	"nericoin/internal/cli"
	"nericoin/internal/nerichain"
)

func main() {
	nc := nerichain.CreateNerichain()
	cli.Execute(nc)

	// Nerichain := nerichain.CreateNerichain()
	// Nerichain.AddNeri("Daddy")
	// Nerichain.AddNeri("Mummy")
	// Nerichain.AddNeri("Uncle")
	// spew.Dump(Nerichain)
}
