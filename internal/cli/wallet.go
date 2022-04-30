package cli

import (
	"fmt"
	"nericoin/internal/nerichain"

	"github.com/spf13/cobra"
)

var (
	walletCmd = &cobra.Command{
		Use:   "wallet [address]",               // cmd name
		Short: "View wallet amount for account", // Description
		Long:  ``,
		Run:   ViewWallet, // Command to be ran
	}
)

func ViewWallet(cmd *cobra.Command, args []string) {
	n := nerichain.CreateNerichain("")
	balance := 0
	UTXOs := n.FindUTXOForAddress(args[0])

	for _, output := range UTXOs {
		balance += output.Value
	}

	fmt.Printf("Balance of '%s':'%d'\n", args[0], balance)
}

func init() {
	rootCmd.AddCommand(walletCmd)
}
