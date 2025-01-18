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
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users, sets, decks, or cards",
	Long: `Makes a single get request to the endpoint associated with the object
type that you are attempting to request`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var listCardCmd = &cobra.Command{
	Use:   "card",
	Short: "List all cards available in the database",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		cards, err := mtgjson.Card.IndexCards() // implement ownership here
		if err != nil {
			fmt.Println("error: Failed to fetch a list of cards (", err.Error(), ")")
			os.Exit(1)
		}

		fmt.Println("Query:")

		for _, card := range *cards {
			fmt.Println("\t", card.ColorIdentity, "-", card.Identifiers.MtgjsonV4Id, "-", card.Name, "(", card.MtgjsonApiMeta.Owner, ")")
		}

		fmt.Println("\n Number of Cards: ", len(*cards))
	},
}

func init() {
	listCmd.PersistentFlags().IntP("limit", "l", 100, "Limit the number of cards returned in the query")
	viper.BindPFlag("api.limit", listCmd.PersistentFlags().Lookup("limit"))

	listCmd.AddCommand(listCardCmd)
	rootCmd.AddCommand(listCmd)
}
