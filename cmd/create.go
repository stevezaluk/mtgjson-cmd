package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Provides an interface for manually creating decks and sets",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("pre-create called")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
