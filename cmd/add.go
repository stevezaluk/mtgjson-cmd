package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a card to a set or deck",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 4 {
			fmt.Println("[error] Invalid number of arguments for add command")
			fmt.Println("[error] You must pass a type, code, uuid (and board if a deck)")
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
