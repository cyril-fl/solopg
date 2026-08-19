package cmd

import (
	"solopg/internal/app"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Launch the game",
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.Start()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
