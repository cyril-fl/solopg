package cmd

import (
	"solopg/internal/app"

	"github.com/spf13/cobra"
)

var tryCmd = &cobra.Command{
	Use:   "try",
	Short: "Launch the game in try mode",
	Run: func(cmd *cobra.Command, args []string) {
		app.Try()
	},
}

func init() {
	rootCmd.AddCommand(tryCmd)
}
