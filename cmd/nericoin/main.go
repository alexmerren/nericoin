package main

import (
	"nericoin/internal/cli"
	"nericoin/internal/nerichain"
)

func main() {
	Nerichain := nerichain.CreateNerichain()
	cli.Execute(Nerichain)
}
