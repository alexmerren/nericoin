package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	newNeriCmd = &cobra.Command{
		Use:   "new-nerichain", // cmd name
		Short: "Make a new Nerichain", // Description
		Long:  ``,
		Args: cobra.MinimumNArgs(1),
		Run:   NewNerichain, // Command to be ran
	}
)

func NewNerichain(cmd *cobra.Command, args []string) {

	fmt.Println("New Neri")
}

func init() {
    rootCmd.AddCommand(newNeriCmd)
}