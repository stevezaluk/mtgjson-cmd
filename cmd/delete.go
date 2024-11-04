package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stevezaluk/mtgjson-sdk-client/deck"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a set, deck, or card based on a code/UUID",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("[error] A type and uuid/code must be provided to the delete command")
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		passedType := args[0]

		if passedType == "card" {
			fmt.Println("[error] Not implemented yet")
			os.Exit(0)
		} else if passedType == "deck" {
			code := args[1]

			if len(code) > 5 {
				fmt.Println("[error] A deck code cannot exceed 5 characters")
				os.Exit(0)
			}

			_, err := deck.DeleteDeck(code)
			if err != nil {
				fmt.Println("[error] ", err.Error())
				os.Exit(0)
			}

			fmt.Println("Successfully deleted deck: ", code)
		} else if passedType == "set" {
			fmt.Println("[error] Not implemented yet")
			os.Exit(0)
		} else {
			fmt.Println("[error] Invalid type for delete command: ", passedType)
			os.Exit(0)
		}

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
