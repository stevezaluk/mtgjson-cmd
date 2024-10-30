package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	card_model "github.com/stevezaluk/mtgjson-models/card"
	"github.com/stevezaluk/mtgjson-models/errors"
	card "github.com/stevezaluk/mtgjson-sdk-client/card"
)

var cardCmd = &cobra.Command{
	Use:   "card",
	Short: "Fetch and display card metadata",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		length := len(args)
		if length == 0 {
			fmt.Println("[error] You must pass a UUID (mtgjsonV4Id) to get a card")
			os.Exit(0)
		} else if length != 1 {
			fmt.Println("[error] You can only pass 1 UUID into the 'card' command")
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		uuid := args[0]
		result, err := card.GetCard(uuid)
		if err == errors.ErrNoCard {
			fmt.Println("[error] No card found with the UUID: ", uuid)
			os.Exit(0)
		} else if err == errors.ErrInvalidUUID {
			fmt.Println("[error] Invalid UUID: ", uuid)
			os.Exit(0)
		}

		DisplayCard(result)
	},
}

func init() {
	rootCmd.AddCommand(cardCmd)
}

func DisplayCard(card card_model.Card) {
	verbose, _ := rootCmd.Flags().GetBool("verbose")

	fmt.Println("Name: ", card.Name)
	fmt.Println("UUID: ", card.Identifiers.MTGJsonV4Id)
	fmt.Println("Set Code: ", card.SetCode)
	fmt.Println("CMC: ", card.ConvertedManaCost)
	if card.ManaCost == "" {
		fmt.Println("Mana Cost: N/A")
	} else {
		fmt.Println("Mana Cost: ", card.ManaCost)
	}
	fmt.Println("Color Identity: ", card.ColorIdentity)
	fmt.Println("Type: ", card.Type)
	if verbose {
		fmt.Println("Types: ", card.Types)
		fmt.Println("Sub-Types: ", card.Subtypes)
		fmt.Println("Super-Types: ", card.Supertypes)
	}
	fmt.Println("Text: ", card.Text)
	fmt.Println("Rarity: ", card.Rarity)

	fmt.Println("\nLegality:")
	fmt.Println("Commander: ", card.Legalities.Commander)
	fmt.Println("Modern: ", card.Legalities.Modern)
	fmt.Println("Standard: ", card.Legalities.Standard)
	fmt.Println("Pauper: ", card.Legalities.Pauper)
	fmt.Println("Vintage: ", card.Legalities.Vintage)
	if verbose {
		fmt.Println("Alchemy: ", card.Legalities.Alchemy)
		fmt.Println("Brawl: ", card.Legalities.Brawl)
		fmt.Println("Duel: ", card.Legalities.Duel)
		fmt.Println("Explorer: ", card.Legalities.Explorer)
		fmt.Println("Future: ", card.Legalities.Future)
		fmt.Println("Gladiator: ", card.Legalities.Explorer)
		fmt.Println("Historic: ", card.Legalities.Historic)
		fmt.Println("Historic Brawl: ", card.Legalities.HistoricBrawl)
		fmt.Println("Legacy: ", card.Legalities.Legacy)
		fmt.Println("Oath Breaker: ", card.Legalities.OathBreaker)
		fmt.Println("Old School: ", card.Legalities.OldSchool)
		fmt.Println("Pauper Commander: ", card.Legalities.PauperCommander)
		fmt.Println("Penny: ", card.Legalities.Penny)
		fmt.Println("Pioneer: ", card.Legalities.Pioneer)
		fmt.Println("Predh: ", card.Legalities.Predh)
		fmt.Println("Pre-Modern: ", card.Legalities.PreModern)
		fmt.Println("Standard Brawl: ", card.Legalities.StandardBrawl)
		fmt.Println("Timeless: ", card.Legalities.Timeless)
	}

	if verbose {
		fmt.Println("\nRulings:")
		for _, value := range card.Rulings {
			fmt.Println(value.Date, value.Text)
		}
	}

}
