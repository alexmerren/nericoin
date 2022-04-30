package cli

import (
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
	//fmt.Println("Sending Nericoin 🪙")
}

func init() {
	rootCmd.AddCommand(sendCmd)
}
