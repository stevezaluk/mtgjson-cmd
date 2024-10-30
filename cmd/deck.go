package cmd

import (
	"github.com/spf13/cobra"
)

var deckCmd = &cobra.Command{
	Use:   "deck",
	Short: "Fetch and display deck metadata",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(deckCmd)
}
