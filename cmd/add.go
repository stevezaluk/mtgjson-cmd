package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	deck_model "github.com/stevezaluk/mtgjson-models/deck"
	"github.com/stevezaluk/mtgjson-models/errors"
	"github.com/stevezaluk/mtgjson-sdk-client/deck"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a card to a set or deck",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			fmt.Println("[error] Invalid number of arguments for add command")
			fmt.Println("[error] You must pass a type, code, uuid (and board if a deck)")
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		passedType := args[0]

		if passedType == "set" {
			fmt.Println("[error] Not implemented yet")
			os.Exit(0)
		} else if passedType == "deck" {
			code := args[1]
			cards := []string{args[2]}

			board, _ := cmd.Flags().GetString("board")
			invalid, noExist, err := deck.AddCards(code, cards, board)
			if err != nil {
				fmt.Println("[error]", err.Error())
				if err == errors.ErrDeckUpdateFailed {

					invalids := map[string][]string{"Invalid Cards:": invalid, "No Exist: ": noExist}
					for key, value := range invalids {
						fmt.Println(key)
						for _, val := range value {
							fmt.Println(val)
						}
					}

				}
			}
		} else {
			fmt.Println("[error] Invalid type for add command. Can be: set, deck")
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().StringP("board", "b", deck_model.MAINBOARD, "Specify the board you want to add to if working with a deck")
}
