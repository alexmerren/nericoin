package cli

import (
	"fmt"
	"nericoin/internal/nerichain"
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
	n := nerichain.CreateNerichain("")
	amountInt, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println(err)
	}
	tx := n.NewUTXOTransaction(args[0], args[1], amountInt)
	n.AddNeri([]*transaction.Transaction{tx})
}

func init() {
	rootCmd.AddCommand(sendCmd)
}
