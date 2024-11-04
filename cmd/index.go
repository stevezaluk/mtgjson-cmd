package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "List all items of a specific type. Can be deck, card, or set",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("[error] A type must be passed to the index command. Can be: deck, card, or set")
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		passedType := args[0]

		if passedType == "deck" {

		} else if passedType == "card" {

		} else if passedType == "set" {
			fmt.Println("[error] Not implemented yet")
			os.Exit(0)
		} else {
			fmt.Println("[error] Invalid type for index command: ", passedType)
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(indexCmd)
}
