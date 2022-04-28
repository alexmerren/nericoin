package cli

import (
	"fmt"
	"nericoin/internal/neri"
	"strconv"

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
	val, _ := strconv.ParseInt(args[2], 10, 64)
	chain.AddNeri(neri.Transaction{
		Ant:   args[0],
		Onio:  args[1],
		Value: val,
	})
}

func init() {
	rootCmd.AddCommand(addNeriCmd)
}
