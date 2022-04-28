package cli

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

var (
	viewNeriCmd = &cobra.Command{
		Use:   "view-nerichain", // cmd name
		Short: "View the Nerichain", // Description
		Long:  ``,
		Run:   ViewNerichain, // Command to be ran
	}
)

func ViewNerichain(cmd *cobra.Command, args []string) {

	fmt.Println("Viewing Nerichain")
	spew.Dump(chain)
}

func init() {
    rootCmd.AddCommand(viewNeriCmd)
}