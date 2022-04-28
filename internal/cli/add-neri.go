package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	addNeriCmd = &cobra.Command{
		Use:   "add-neri", // cmd name
		Short: "Add Neri to the Nerichain", // Description
		Long:  ``,
		Run:   AddNeriToNerichain, // Command to be ran
	}
)

func AddNeriToNerichain(cmd *cobra.Command, args []string) {

	fmt.Println("Adding Neri to Nerichain...")

	if len(args) == 0 {
		fmt.Println("Error: No Neri hash")
		return
	}
	neri_hash := args[0]
	fmt.Println("Creating block with hash " + neri_hash)
}

func init() {
    rootCmd.AddCommand(addNeriCmd)
}