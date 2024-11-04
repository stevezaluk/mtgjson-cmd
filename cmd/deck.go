package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stevezaluk/mtgjson-models/card"
	deck_model "github.com/stevezaluk/mtgjson-models/deck"
	"github.com/stevezaluk/mtgjson-models/errors"
	deck "github.com/stevezaluk/mtgjson-sdk-client/deck"
)

var deckCmd = &cobra.Command{
	Use:   "deck",
	Short: "Fetch and display deck metadata",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		length := len(args)
		if length == 0 {
			fmt.Println("[error] You must pass a deck code to get a deck")
			os.Exit(0)
		} else if length != 1 {
			fmt.Println("[error] You can only pass 1 deck code into the 'deck' command")
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		code := args[0]
		result, err := deck.GetDeck(code)
		if err == errors.ErrNoDeck {
			fmt.Println("[error] No deck found with the code: ", code)
			os.Exit(0)
		}

		DisplayDeck(result)
	},
}

func init() {
	rootCmd.AddCommand(deckCmd)
}

func DisplayDeck(deck deck_model.Deck) {
	verbose, _ := rootCmd.Flags().GetBool("verbose")

	fmt.Println("Name: ", deck.Name)
	fmt.Println("Code: ", deck.Code)
	fmt.Println("Type: ", deck.Type)
	fmt.Println("Release Date: ", deck.ReleaseDate)

	if verbose {
		boards := map[string][]card.Card{
			"Mainboard": deck.Contents.Mainboard,
			"Sideboard": deck.Contents.Sideboard,
			"Commander": deck.Contents.Sideboard,
		}

		for key, value := range boards {
			fmt.Println("\n|-", key)
			for _, card := range value {
				fmt.Printf("|---[%s] %s\n", card.Identifiers.MTGJsonV4Id, card.Name)
			}
		}
	}
}
