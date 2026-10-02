package cmd

import (
	"fmt"
	"solopg/config"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     cmd.Version.Use,
	Short:   cmd.Version.Short,
	Long:    cmd.Version.Long,
	Example: cmd.Version.Example,
	Run: func(cobra *cobra.Command, args []string) {
		fmt.Println(config.String())
	},
}

func init() {
	cmdRoot.AddCommand(versionCmd)
}
