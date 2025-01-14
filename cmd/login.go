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
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "",
	Long:  `Exchange user credentials for an OAuth Access Token`,
	PreRun: func(cmd *cobra.Command, args []string) {
		email := viper.GetString("api.email")
		if email == "" {
			fmt.Println("error: a email must be passed to make an authentication request")
			os.Exit(1)
		}

		password := viper.GetString("api.password")
		if password == "" {
			fmt.Println("error: a password must be passed to make an authentication request")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := mtgjson.Auth.SetAuthToken(viper.GetString("api.email"), viper.GetString("api.password"))
		if err != nil {
			fmt.Println("error: Failed to fetch access token (", err.Error(), ")")
			os.Exit(1)
		}

		err = viper.WriteConfig()
		if err != nil {
			fmt.Println("error: Failed to write config file with access token (", err.Error(), ")")
		}

		fmt.Println("Fetched access token for: ", viper.GetString("api.email"))
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
