package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stevezaluk/mtgjson-sdk-client/health"
	"os"
)

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Make a call to the /health endpoint to determine the health of the server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := health.GetHealth()
		if err != nil {
			fmt.Println("[error]", err.Error())
			os.Exit(0)
		}

		if result {
			fmt.Println("[success] Server is healthy")
		} else {
			fmt.Println("[error] Server is not healthy or offline")
		}
	},
}

func init() {
	rootCmd.AddCommand(healthCmd)
}
