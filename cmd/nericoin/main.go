package main

import (
	"nericoin/internal/nerichain"
)

func main() {
	Nerichain := nerichain.CreateNerichain()
	//Nerichain.AddNeri("Daddy")
	//Nerichain.AddNeri("Mummy")
	//Nerichain.AddNeri("Uncle")
	Nerichain.ViewNerichain()
}
