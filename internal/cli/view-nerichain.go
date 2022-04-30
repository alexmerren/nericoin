package cli

import (
	"nericoin/internal/nerichain"

	"github.com/spf13/cobra"
)

var (
	viewNeriCmd = &cobra.Command{
		Use:   "view-nerichain",     // cmd name
		Short: "View the Nerichain", // Description
		Long:  ``,
		Run:   ViewNerichain, // Command to be ran
	}
)

func ViewNerichain(cmd *cobra.Command, args []string) {
	n := nerichain.CreateNerichain("")
	n.ViewNerichain()
}

func init() {
	rootCmd.AddCommand(viewNeriCmd)
}
