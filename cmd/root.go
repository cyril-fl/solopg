package cmd

import (
	"fmt"
	"os"
	"solopg/config"

	"github.com/spf13/cobra"
)

var c = config.Current
var cmd = c.Commands

var rootCmd = &cobra.Command{
	Use:     c.Name,
	Short:   cmd.Root.Short,
	Long:    cmd.Root.Long,
	Version: config.Version,
	Example: cmd.Root.Example,

	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
