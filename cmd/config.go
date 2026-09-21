package cmd

import (
	"solopg/app/utils/log"
	"solopg/config"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     cmd.Config.Use,
	Short:   cmd.Config.Short,
	Long:    cmd.Config.Long,
	Example: cmd.Config.Example,
	RunE: func(cmd *cobra.Command, args []string) error {
		log.ParseJson(config.Current)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
