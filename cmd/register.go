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
	"os"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "",
	Long:  `Register a new user account for the MTGJSON API`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			fmt.Println("error: Three arguments must be passed to this command: username, email and password")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		email := args[1]
		password := args[2]

		message, err := mtgjson.Auth.RegisterUser(email, username, password)
		if err != nil {
			fmt.Println("error: Failed to register user (", err.Error(), ")")
			os.Exit(1)
		}

		fmt.Println("Successfully Registered User: ", message.Message)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
