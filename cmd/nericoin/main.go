package main

import (
	"nericoin/internal/cli"
	"nericoin/internal/nerichain"
)

func main() {


	// Nerichain := nerichain.CreateNerichain()
	// Nerichain.AddNeri("Daddy")
	// Nerichain.AddNeri("Mummy")
	// Nerichain.AddNeri("Uncle")
	// spew.Dump(Nerichain)
	Nerichain := nerichain.CreateNerichain()
	cli.Execute(Nerichain)
	//Nerichain.AddNeri("Daddy")
	//Nerichain.AddNeri("Mummy")
	//Nerichain.AddNeri("Uncle")

}
