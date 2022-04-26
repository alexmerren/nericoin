package main

import (
	"fmt"
	"nericoin/internal/nerichain"
)

func main() {
	fmt.Println("Hello, world!")
	Nerichain := nerichain.NeriChain{}
	TheOneAndOnlyNeri := nerichain.CreateGenesisBlock()
	Nerichain = append(Nerichain, TheOneAndOnlyNeri)

	verified_block := TheOneAndOnlyNeri.VerifyNeri()
	fmt.Println(verified_block)
}
