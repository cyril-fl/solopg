package cmd

import (
	"solopg/app/shared/utils/debug"
	"solopg/config"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     cmd.Config.Use,
	Short:   cmd.Config.Short,
	Long:    cmd.Config.Long,
	Example: cmd.Config.Example,
	RunE: func(cobra *cobra.Command, args []string) error {
		debug.ParseJson(config.Current)

		return nil
	},
}

func init() {
	cmdRoot.AddCommand(configCmd)
}
