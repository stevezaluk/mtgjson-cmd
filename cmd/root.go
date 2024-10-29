package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "mtgjson-cmd",
	Short: "A command line application for managing your MTGJSON API instance",
	Long:  ``,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("verbose", "v", false, "Increase verbosity across commands")
	rootCmd.Flags().StringP("config", "c", "~/.config/mtgjson-cmd/config.json", "Override the default config file")
}
