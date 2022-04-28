package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	addNeriCmd = &cobra.Command{
		Use:   "add-neri",                  // cmd name
		Short: "Add Neri to the Nerichain", // Description
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run:   AddNeriToNerichain, // Command to be ran
	}
)

func AddNeriToNerichain(cmd *cobra.Command, args []string) {

	fmt.Println("Adding Neri to Nerichain...")
}

func init() {
	rootCmd.AddCommand(addNeriCmd)
}
