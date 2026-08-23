package cmd

import (
	"solopg/internal/app"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     cmd.Run.Use,
	Short:   cmd.Run.Short,
	Long:    cmd.Run.Long,
	Example: cmd.Run.Example,
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.Start()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
