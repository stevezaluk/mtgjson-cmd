package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Make a call to the /health endpoint to determine the health of the server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("health called")
	},
}

func init() {
	rootCmd.AddCommand(healthCmd)
}
