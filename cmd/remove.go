package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	deck_model "github.com/stevezaluk/mtgjson-models/deck"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a card from a deck or a set",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("pre-remove called")
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.PersistentFlags().StringP("board", "b", deck_model.MAINBOARD, "Specify the board you want to add to if working with a deck")
}
