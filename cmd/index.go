package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "List all items of a specific type. Can be deck, card, or set",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("index called")
	},
}

func init() {
	rootCmd.AddCommand(indexCmd)
}
