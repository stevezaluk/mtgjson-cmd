package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stevezaluk/mtgjson-sdk-client/config"
	"github.com/stevezaluk/mtgjson-sdk-client/context"
)

var rootCmd = &cobra.Command{
	Use:   "mtgjson-cmd",
	Short: "A command line application for managing your MTGJSON API instance",
	Long:  ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config_path, _ := cmd.Flags().GetString("config")

		config := config.ParseConfig(config_path)
		context.InitConfig(config)
		context.InitUri(config)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Increase verbosity across commands")
	rootCmd.PersistentFlags().StringP("config", "c", "~/.config/mtgjson-cmd/config.json", "Override the default config file")
}
