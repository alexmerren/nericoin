package cli

import (
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
	chain.ViewNerichain()
}

func init() {
	rootCmd.AddCommand(viewNeriCmd)
}
