package cli

import (
	"nericoin/internal/nerichain"

	"github.com/spf13/cobra"
)

var (
	createNerichainCmd = &cobra.Command{
		Use:   "create-nerichain [address]",                  // cmd name
		Short: "create a nerichain with your wallet address", // Description
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run:   CreateNerichain, // Command to be ran
	}
)

func CreateNerichain(cmd *cobra.Command, args []string) {
	nerichain.CreateNerichain(args[0])
}

func init() {
	rootCmd.AddCommand(createNerichainCmd)
}
