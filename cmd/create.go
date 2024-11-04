package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	deck "github.com/stevezaluk/mtgjson-sdk-client/deck"
	"os"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Provides an interface for manually creating decks and sets",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("[error] A type must be passed to the created command. Can be: deck, set")
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		passedType := args[0]

		if passedType == "deck" {
			if len(args) != 4 {
				fmt.Println("[error] A name, deck code, and deck type must be passed to the create command")
				os.Exit(0)
			}

			name := args[1]
			code := args[2]
			deckType := args[3]

			result, err := deck.CreateDeck(name, code, deckType)
			if !result {
				fmt.Println("[error]", err.Error())
			}
		} else if passedType == "set" {
			fmt.Println("[error] Creating sets is not implemented yet")
			os.Exit(0)
		} else {
			fmt.Println("[error] Invalid type for create command: ", passedType)
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().StringP("import", "i", "", "Pass a .json file to use for creation. Should define the type you want to import")
}
