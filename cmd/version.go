package cmd

import (
	"fmt"
	"solopg/internal/infrastructure/config"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     cmd.Version.Use,
	Short:   cmd.Version.Short,
	Long:    cmd.Version.Long,
	Example: cmd.Version.Example,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.String())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
