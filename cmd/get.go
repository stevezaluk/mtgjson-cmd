/*
Copyright © 2025 Steven Zaluk

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Fetch metadata for a user, deck, set, or card",
	Long: `Makes a single GET request to the endpoint associated with the object type you are
trying to request. In most cases an owner (--owner) will need to passed. This will default to the system
owner if one is not specified`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var getCardCmd = &cobra.Command{
	Use:   "card",
	Short: "Fetch metadata for a single card",
	Long:  "",
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("error: A UUID must be passed to fetch a card")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// check for args length here
		owner := viper.GetString("api.owner")

		card, err := mtgjson.Card.GetCard(args[0], owner)
		if err != nil {
			fmt.Println("error: Failed to fetch card (", err.Error(), ")")
			os.Exit(1)
		}

		fmt.Println("Name:", card.Name)
		fmt.Println("Set Code:", card.SetCode)
		fmt.Println("Mana Cost: ", card.ManaCost)
		fmt.Println("CMC (Converted Mana Cost): ", card.ConvertedManaCost)
		fmt.Println("Color Identity: ", card.ColorIdentity)
		fmt.Println("Type:", card.Type)
		fmt.Println("Sub-Type:", card.Subtypes)
		fmt.Println("Description:", card.Text)
		if strings.Contains(card.Type, "Creature") {
			fmt.Println("Creature Data:")
			fmt.Println("\t Power:", card.Power)
			fmt.Println("\tToughness:", card.Toughness)
		}
		fmt.Println("Legalities:")
		fmt.Println("\tCommander:", card.Legalities.Commander)
		fmt.Println("\tStandard: ", card.Legalities.Standard)
		fmt.Println("\tModern:", card.Legalities.Modern)
		fmt.Println("\tVintage:", card.Legalities.Vintage)
		fmt.Println("\tLegacy:", card.Legalities.Legacy)
		if viper.GetBool("verbose") {
			fmt.Println("\tStandard Brawl:", card.Legalities.Standardbrawl)
			fmt.Println("\tBrawl:", card.Legalities.Brawl)
			fmt.Println("\tAlchemey:", card.Legalities.Alchemy)
			fmt.Println("\tDuel:", card.Legalities.Duel)
			fmt.Println("\tExplorer:", card.Legalities.Explorer)
			fmt.Println("\tFuture:", card.Legalities.Future)
			fmt.Println("\tGladiator:", card.Legalities.Gladiator)
			fmt.Println("\tHistoric:", card.Legalities.Historic)
			fmt.Println("\tHistoric Brawl: ", card.Legalities.HistoricBrawl)
			fmt.Println("\tOathbreaker:", card.Legalities.Oathbreaker)
			fmt.Println("\tOld School:", card.Legalities.Oldschool)
			fmt.Println("\tPenny:", card.Legalities.Penny)
			fmt.Println("\tPauper:", card.Legalities.Pauper)
			fmt.Println("\tPauper Commander: ", card.Legalities.Paupercommander)
			fmt.Println("\tPredh:", card.Legalities.Predh)
			fmt.Println("\tTimeless:", card.Legalities.Timeless)

			if len(card.Rulings) != 0 {
				fmt.Println("Rulings:")
				for _, ruling := range card.Rulings {
					fmt.Println("\t(", ruling.Date, ")", " - ", ruling.Text)
				}
			}

			fmt.Println("Identifiers:")
			fmt.Println("\tMTGJSON V4 ID:", card.Identifiers.MtgjsonV4Id)
			fmt.Println("\tCardKingdom ID:", card.Identifiers.CardKingdomId)
			fmt.Println("\tScryfall ID:", card.Identifiers.ScryfallId)
			fmt.Println("\tTCG Player Product ID:", card.Identifiers.TcgplayerProductId)
			fmt.Println("\tCard Sphere ID:", card.Identifiers.CardsphereId)

			fmt.Println("Artist Information:")
			fmt.Println("\tArtist:", card.Artist)
			fmt.Println("\tArtist IDs:")
			for _, artistId := range card.ArtistIds {
				fmt.Println("\t\t", artistId)
			}

			fmt.Println("Boolean Data:")
			fmt.Println("\tIs Funny: ", card.IsFunny)
			fmt.Println("\tIs Full Art: ", card.IsFullArt)
			fmt.Println("\tIs Oversized: ", card.IsOversized)
			fmt.Println("\tIs Online Only: ", card.IsOnlineOnly)
			fmt.Println("\tIs Promo: ", card.IsPromo)
			fmt.Println("\tIs Alternative: ", card.IsAlternative)
			fmt.Println("\tIs Rebalanced: ", card.IsRebalanced)
			fmt.Println("\tIs Reprint: ", card.IsReprint)
			fmt.Println("\tIs Reserved: ", card.IsReserved)
			fmt.Println("\tIs Starter: ", card.IsStarter)
			fmt.Println("\tIs Textless: ", card.IsTextless)
			fmt.Println("\tIs Timeshifted: ", card.IsTimeshifted)
			fmt.Println("\tIs Story Spotlight:", card.IsStorySpotlight)

			fmt.Println("Variations:")
			for _, variation := range card.Variations {
				fmt.Println("\t", variation)
			}

			// card related bools

			fmt.Println("MTGJSON API Metadata:")
			fmt.Println("\tOwner:", card.MtgjsonApiMeta.Owner)
			fmt.Println("\tType:", card.MtgjsonApiMeta.Type)
			fmt.Println("\tSub-Type:", card.MtgjsonApiMeta.Subtype)
			fmt.Println("\tCreation Date:", card.MtgjsonApiMeta.CreationDate)
			fmt.Println("\tModified Date:", card.MtgjsonApiMeta.ModifiedDate)
		}

	},
}

func init() {

	getCmd.PersistentFlags().String("owner", "system", "Set the owner to use for the query. Defaults to the email provided in the config file")
	viper.BindPFlag("api.owner", getCmd.PersistentFlags().Lookup("owner"))

	getCmd.AddCommand(getCardCmd)
	rootCmd.AddCommand(getCmd)
}
