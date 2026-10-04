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
	Run: func(cobra *cobra.Command, args []string) {
		debug.ParseJson(config.Current)
	},
}

func init() {
	cmdRoot.AddCommand(configCmd)
}
