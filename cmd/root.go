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
	"github.com/stevezaluk/mtgjson-sdk-client/api"
	"github.com/stevezaluk/mtgjson-sdk-client/config"
	"os"
)

const (
	defaultConfigPath = "/.config/mtgjson-cmd"
	defaultConfigName = "config.json"
)

var cfgFile string
var mtgjson *api.MtgjsonApi

var rootCmd = &cobra.Command{
	Use:   "mtgjson-cmd",
	Short: "A command line application for interacting with the MTGJSON API",
	Long:  `Provides fully featured access to the API using the command line`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initApi)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mtgjson-cmd.yaml)")

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbosity across all commands")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

func initConfig() {
	err := config.ReadConfigFile(cfgFile)
	if err != nil {
		fmt.Println("error: Failed to read config file (", err.Error(), ")")
		os.Exit(1)
	}
}

func initApi() {
	mtgjson = api.New()
}
