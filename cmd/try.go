package cmd

import (
	src "solopg/app"

	"github.com/spf13/cobra"
)

var tryCmd = &cobra.Command{
	Use:     cmd.Try.Use,
	Short:   cmd.Try.Short,
	Long:    cmd.Try.Long,
	Example: cmd.Try.Example,
	RunE: func(cmd *cobra.Command, args []string) error {
		return src.Try()
	},
}

func init() {
	rootCmd.AddCommand(tryCmd)
}
