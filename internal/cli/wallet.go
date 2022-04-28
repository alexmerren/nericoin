package cli

import (
	"fmt"
	"nericoin/internal/wallet"
	"strconv"

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
	fmt.Println("In wallet: '" + args[0] + "' you have: \n" + strconv.FormatInt(wallet.WalletSum(chain, args[0]), 10))
}

func init() {
	rootCmd.AddCommand(walletCmd)
}
