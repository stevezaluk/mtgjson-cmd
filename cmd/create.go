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
	"github.com/gofrs/uuid/v5"

	"github.com/spf13/viper"
	"github.com/stevezaluk/mtgjson-models/card"
	"github.com/stevezaluk/mtgjson-models/meta"
	"os"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an Card, Deck, or a Set",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var createCardCmd = &cobra.Command{
	Use:   "card",
	Short: "Create a new card",
	Long: `At a minimum a card name, and an ID must be passed. If an owner is not passed, one will be
created under the logged in user's account'`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("error: 1 argument is required to create a card. A name must be passed")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var mtgjsonv4id string
		if len(args) == 1 {
			fmt.Println("warn: A MTGJSONv4ID was not passed for the card. One will be generated for you")
			namespace, err := uuid.NewV4()
			if err != nil {
				fmt.Println("error: Failed to generate UUID for card (", err.Error(), ")")
				os.Exit(1)
			}

			mtgjsonv4id = uuid.NewV5(namespace, args[0]).String()
		} else {
			mtgjsonv4id = args[1]
		}

		newCard := &card.CardSet{
			Name: args[0],
			Identifiers: &meta.CardIdentifiers{
				MtgjsonV4Id: mtgjsonv4id,
			},
		}

		owner, _ := cmd.Flags().GetString("owner")
		message, err := mtgjson.Card.NewCard(newCard, owner)
		if err != nil {
			fmt.Println("error: Failed to create card (", err.Error(), ")")
			fmt.Println("Message from API: ", message.Message)
			os.Exit(1)
		}

		fmt.Println("Created:", args[0])
	},
}

func init() {
	createCmd.PersistentFlags().String("owner", "system", "Set the owner to use for the query. Defaults to the email provided in the config file")
	viper.BindPFlag("api.owner", createCmd.PersistentFlags().Lookup("owner"))

	createCmd.AddCommand(createCardCmd)
	rootCmd.AddCommand(createCmd)
}
