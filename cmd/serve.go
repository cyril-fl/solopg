package cmd

import (
	src "solopg/app/cmdserve"
	"solopg/app/shared/utils/cmdargs"

	"github.com/spf13/cobra"
)

var cmdServe = &cobra.Command{
	Use:     cmd.Serve.Use,
	Short:   cmd.Serve.Short,
	Long:    cmd.Serve.Long,
	Example: cmd.Serve.Example,
	RunE: func(cobra *cobra.Command, args []string) error {
		return src.RunCmdServe(cobra.Flags)
	},
}

func init() {
	cmdargs.MakeFlags(cmdServe, cmd.Serve.Args)

	cmdRoot.AddCommand(cmdServe)
}
