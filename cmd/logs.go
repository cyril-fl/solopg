package cmd

import (
	"solopg/app/cmdlogs"
	"solopg/app/shared/utils/cmdargs"

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
	cmdargs.MakeFlags(cmdLogs, cmd.Logs.Args)

	cmdRoot.AddCommand(cmdLogs)
}
