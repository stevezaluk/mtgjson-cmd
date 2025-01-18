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

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a user, card, deck, or set",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var deleteCardCmd = &cobra.Command{
	Use:   "card",
	Short: "Delete a card using its MTGJSON V4 ID",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("error: A UUID must be passed to delete a card")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		owner, _ := cmd.Flags().GetString("owner")
		fmt.Println(owner)

		message, err := mtgjson.Card.DeleteCard(args[0], owner)
		if err != nil {
			fmt.Println("error: Failed to delete card (", err.Error(), ")")
			if viper.GetBool("verbose") {
				fmt.Println("Message from API: ", message.Message)
			}
			os.Exit(1)
		}

		fmt.Println("Deleted: ", args[0])
	},
}

func init() {
	deleteCmd.PersistentFlags().String("owner", "system", "Set the owner to use for the query. Defaults to the email provided in the config file")
	viper.BindPFlag("api.owner", deleteCmd.PersistentFlags().Lookup("owner"))

	deleteCmd.AddCommand(deleteCardCmd)
	rootCmd.AddCommand(deleteCmd)
}
