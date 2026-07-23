package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Launch the game",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("running game...")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
