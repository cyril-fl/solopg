package cmd

import (
	"solopg/internal/app"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Launch the game",
	RunE: func(cmd *cobra.Command, args []string) error {
		app.Start()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
