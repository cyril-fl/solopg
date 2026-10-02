package cmd

import (
	"solopg/app/cmdlogs"

	"github.com/spf13/cobra"
)

var cmdLogs = &cobra.Command{
	Use:     cmd.Logs.Use,
	Short:   cmd.Logs.Short,
	Long:    cmd.Logs.Long,
	Example: cmd.Logs.Example,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmdlogs.RunCmdLogs(cmd.Flags)
	},
}

func init() {
	cmdLogs.Flags().IntP("tail", "t", 0, "Show only the last N log entries")
	cmdLogs.Flags().StringP("filter", "f", "", "Show only logs at the specified level")

	cmdRoot.AddCommand(cmdLogs)
}
