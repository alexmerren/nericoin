package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Root command for the CLI (nericoin)
var rootCmd = &cobra.Command{
	Use:   "nericoin",
	Short: "All the CLI for Nericoin",
	Long: ``,
  }

  func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
  }


// func CreateNerichain() *nc.Nerichain {
// 	return &nerichain.Nerichain{}
// }

// func ViewNerichain(chain *nc.Nerichain) {}

// func AddNeriToNerichain(data string) *nc.Nerichain {
// 	return &nc.Nerichain{}
// }

// // KILL ME
// // KILL ME
// // KILL ME

// func CreateWallet() {}

// func ViewWalletBalance() {}

// func CreateTransaction() {}
