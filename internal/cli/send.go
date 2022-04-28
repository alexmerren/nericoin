package cli

import (
	"fmt"
	"nericoin/internal/transaction"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	sendCmd = &cobra.Command{
		Use:   "send [sender] [receiver] [amount]",                // cmd name
		Short: "Send Nericoin from your wallet to another wallet", // Description
		Long:  ``,
		Args:  cobra.MinimumNArgs(3),
		Run:   Send, // Command to be ran
	}
)

func Send(cmd *cobra.Command, args []string) {
	fmt.Println("Sending Nericoin 🪙")
	val, _ := strconv.ParseInt(args[2], 10, 64)
	chain.AddNeri(transaction.Transaction{
		Ant:   args[0],
		Onio:  args[1],
		Value: val,
	})
}

func init() {
	rootCmd.AddCommand(sendCmd)
}
