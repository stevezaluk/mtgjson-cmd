package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	deck_model "github.com/stevezaluk/mtgjson-models/deck"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a card from a deck or a set",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			fmt.Println("[error] Invalid number of arguments for add command")
			fmt.Println("[error] You must pass a type, code, uuid (and board if a deck)")
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.PersistentFlags().StringP("board", "b", deck_model.MAINBOARD, "Specify the board you want to add to if working with a deck")
}
