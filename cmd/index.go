package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stevezaluk/mtgjson-sdk-client/card"
	"github.com/stevezaluk/mtgjson-sdk-client/deck"
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
		limit, _ := cmd.Flags().GetInt("limit")

		passedType := args[0]
		if passedType == "deck" {
			results, err := deck.IndexDecks(limit)
			if err != nil {
				fmt.Println("[error]", err.Error())
				os.Exit(0)
			}

			fmt.Printf("Found %d decks\n\n", len(results))
			for _, value := range results {
				fmt.Printf("[%s] %s - %s\n", value.Code, value.Name, value.ReleaseDate)
			}
		} else if passedType == "card" {
			results, err := card.IndexCards(limit)
			if err != nil {
				fmt.Println("[error]", err.Error())
			}

			fmt.Printf("Found %d cards\n\n", len(results))
			for _, value := range results {
				fmt.Printf("[%s] %s - %s\n", value.Identifiers.MTGJsonV4Id, value.Name, value.ColorIdentity)
			}
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

	indexCmd.PersistentFlags().IntP("limit", "l", 100, "Limit the number of items returned")
}
